package users

import "github.com/google/uuid"

type PhoneNumber struct {
	PhoneNumberType string
	Number          string
	ID              uuid.UUID
}
