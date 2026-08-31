package entities

import "github.com/google/uuid"

type User struct {
	UUID           uuid.UUID
	Username       string
	HashedPassword string
	Email          string
}

func NewUser(username string, hashedPassword string, email string) *User {
	return &User{
		Username:       username,
		HashedPassword: hashedPassword,
		Email:          email,
	}
}
