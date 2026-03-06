package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	JobTitle  string
	Salary    int
	JoinedAt  time.Time
}

// Value receiver method
func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

// Pointer receiver method
func (e *Employee) SetSalary(salary int) {
	e.Salary = salary
}

func (e *Employee) SetJoinedAt(t time.Time) {
	e.JoinedAt = t
}

func NewEmployee(id int, firstName, lastName, jobTitle string) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		JobTitle:  jobTitle,
		JoinedAt:  time.Now(),
	}
}

func main() {
	bob := Employee{
		ID:        1,
		FirstName: "Bob",
		LastName:  "Bob",
		JobTitle:  "Software developer",
	}

	fmt.Println(bob.FirstName)
	bob.FullName()

	bob.SetSalary(12500000)
	fmt.Println(bob.Salary)
	bob.SetJoinedAt(time.Now().Add(10 * time.Hour))
	fmt.Println(bob.JoinedAt)
}
