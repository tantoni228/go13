package models

type Member struct {
	UserId string `db:"user_id"`
	RoleId int    `db:"role_id"`
}
