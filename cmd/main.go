package main

import (
	"github.com/irfan44/go-h8-boilerplate/pkg/database"
	"log"

	"github.com/irfan44/go-h8-boilerplate/internal/config"
	"github.com/irfan44/go-h8-boilerplate/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %s\n", err.Error())
	}
}

func main() {
	cfg := config.NewConfig()

	db, err := database.InitPGDB(cfg)
	if err != nil {
		return
	}

	s := server.NewServer(cfg, db)

	s.Run()
}
