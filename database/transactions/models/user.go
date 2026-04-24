package models

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type Profile struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Bio       string `json:"bio"`
	CreatedAt string `json:"created_at"`
}
