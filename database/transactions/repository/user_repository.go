package repository

import (
	"context"
	"database/sql"
	"learning-go/database/transactions/models"

	"golang.org/x/crypto/bcrypt"
)

var insertUserQuery = `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`
var getUsersQuery = `SELECT * FROM users`

var getUserByEmailQuery = `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?`

var insertProfileQuery = `INSERT INTO profiles (user_id, bio) VALUES (?, ?)`

var getProfileByUserIdQuery = `SELECT id, user_id, bio, created_at FROM profiles WHERE user_id = ?`

type UserRepository interface {
	CreateUser(name, email, password string) (int64, error)
	GetUsers() ([]models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetProfileByUserId(userId int) (*models.Profile, error)
}

type SqlUserRepository struct {
	db *sql.DB
}

func NewSqlRepository(db *sql.DB) UserRepository {
	return &SqlUserRepository{db: db}
}

func (s SqlUserRepository) CreateUser(name, email, password string) (int64, error) {
	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	stmt, err := tx.PrepareContext(ctx, insertUserQuery)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmt.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(name, email, hashedPassword)
	if err != nil {
		return 0, err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	profileStmt, err := tx.PrepareContext(ctx, insertProfileQuery)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer profileStmt.Close()
	_, err = profileStmt.Exec(userId, "Hello, I'm "+name)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (s SqlUserRepository) GetUsers() ([]models.User, error) {
	rows, err := s.db.Query(getUsersQuery)
	if err != nil {
		return []models.User{}, nil
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return []models.User{}, nil
		}
		users = append(users, user)
	}
	return users, nil
}

func (s SqlUserRepository) GetUserByEmail(email string) (*models.User, error) {
	row := s.db.QueryRow(getUserByEmailQuery, email)
	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s SqlUserRepository) GetProfileByUserId(userId int) (*models.Profile, error) {
	row := s.db.QueryRow(getProfileByUserIdQuery, userId)
	var profile models.Profile
	err := row.Scan(&profile.Id, &profile.UserId, &profile.Bio, &profile.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
