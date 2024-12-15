package models

type User struct {
	UserId   string   `db:"user_id"`
	Username string   `db:"user_name"`
	Email    Email    `db:"user_email"`
	Password Password `db:"user_password"`
	Bio      string   `db:"user_bio"`
}

type Email string
type Password string

type Users struct {
	Users []User
}
