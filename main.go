package main

import (
	"github.com/joho/godotenv"
	"github.com/sharpvik/log-go"

	"github.com/sharpvik/mess/server/configs"
	"github.com/sharpvik/mess/server/database"
	"github.com/sharpvik/mess/server/server"
)

// init simply loads environment from the .env file.
func init() {
	log.Debug("reading .env ...")
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to read .env")
	}
}

// mustInit is responsible for the primary and the most essential initialization
// code that has to run properly.
func mustInit() (config configs.Config, db *database.Database) {
	config = configs.MustInit()
	db = database.MustInit(config.Database)
	return
}

func main() {
	config, _ := mustInit()
	log.Debug("init successfull")

	serv := server.NewServer(config.Server)
	done := make(chan bool, 1)
	go serv.ServeWithGrace(done)

	<-done
	log.Debug("server stopped")
}
