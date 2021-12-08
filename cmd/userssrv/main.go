package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"users/internal/data"
	"users/internal/routes"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	var parser = argparse.NewParser("nwCraft", "Provides shopping list for desired craft")
	debugLevel := parser.Selector("d", "debug-level", []string{"INFO", "DEBUG"}, &argparse.Options{Required: false, Help: "Logging debug level"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	switch *debugLevel {
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	log.Info().Msg("Starting user startup")
	hostname := os.Getenv(data.DatabaseHostname)
	database := os.Getenv(data.DatabaseName)
	user := os.Getenv(data.DatabaseUser)
	password := os.Getenv(data.DatabasePassword)
	data.Initializer(hostname, database, user, password)
	routes.Routes()
}
