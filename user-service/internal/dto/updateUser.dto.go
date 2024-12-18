package dto

type UpdateUserInput struct {
	Username       OptString
	Email          OptString
	HashedPassword OptString
	Bio            OptString
}
