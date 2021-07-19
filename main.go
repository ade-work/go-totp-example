package main

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"go-totp-example/http"
	"log"
)

const appPort = ":3000"

func main() {
	initEnv()

	app := http.NewRouter()
	log.Fatal(app.Listen(appPort))
}

func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(errors.Wrap(err, "env load error"))
	}
}
