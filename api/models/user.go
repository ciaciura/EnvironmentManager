// models/user.go
package models

type Role string

const (
	Admin    Role = "Admin"
	Employee Role = "Employee"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}
