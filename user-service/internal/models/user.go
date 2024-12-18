package models

type User struct {
	Id             string `db:"id"`
	Username       string `db:"username"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	Bio            string `db:"bio"`
}
