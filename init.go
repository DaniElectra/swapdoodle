package main

import (
	"github.com/PretendoNetwork/swapdoodle/database"
	"github.com/PretendoNetwork/plogger-go"
	"github.com/joho/godotenv"
)

var logger = plogger.NewLogger()

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Warning("Error loading .env file")
	}

	database.ConnectAll()
}
