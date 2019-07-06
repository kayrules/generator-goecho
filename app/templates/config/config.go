package config

import (
	"log"
	"os"
)

// var public settings
var (
	Env    string
	DbHost string
	DbName string
	Port   string
)

func init() {

	Env = os.Getenv("ENV")
	Port = os.Getenv("PORT")

	if Env == "" {
		log.Fatal("cannot find ENV from Env")
	}

	if Port == "" {
		log.Fatal("cannot find PORT from Env")
	}
}
