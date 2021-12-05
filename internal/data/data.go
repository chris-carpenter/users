package data

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"users/internal/users"
)

func SaveUser(user users.User) error {
	fmt.Println(user)
	//TODO: Implement saving a user
	return nil
}

func GetUser(id uuid.UUID) (users.User, error) {
	var user users.User
	//TODO: Implement getting a user
	return user, nil
}
