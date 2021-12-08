package data

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"users/internal/users"
)

type UserDatabase interface {
	SaveUser() error
	GetUser() (users.User, error)
	Initializer()
}

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

func Initializer(hostname string, database string, user string, password string) {
	dbHostname := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, password, hostname, database)
	log.Debug().Str("dbHostName", dbHostname)
	m, err := migrate.New("file://db/migrations", dbHostname)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create mew migrate")
	}
	if err := m.Up(); err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate database")
	}
	log.Info().Str("event", "init_db").Msg("Successfully initialized db")
}
