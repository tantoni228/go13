package models

type User struct {
	UserId string `db:"user_id"`
	Username string `db:"user_name"`
	Email string `db:"user_email"`
	Password string `db:"user_password"`
	Bio string `db:"user_bio"`
}

type Users struct {
	Users []User
}