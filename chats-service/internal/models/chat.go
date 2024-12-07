package models

type Chat struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
