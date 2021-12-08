package users

import (
	"encoding/json"
	"github.com/creasty/defaults"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type User struct {
	Name         string
	Password     string
	Group        string
	ID           uuid.UUID
	Emails       []Email
	PhoneNumbers []PhoneNumber
}

func (p User) String() string {
	val, err := json.Marshal(p)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal person to json")
	}
	return string(val)
}

func (p *User) SetDefaults() {
	if defaults.CanUpdate(p.ID) {
		p.ID = uuid.New()
	}
}

func NewPerson() (User, error) {
	person := &User{}
	var err error
	if err = defaults.Set(person); err != nil {
		log.Error().Err(err).Msg("Error creating empty person")
	}
	return *person, err
}
