package users

import "github.com/google/uuid"

type Email struct {
	EmailType string
	Address   string
	ID        uuid.UUID
}
