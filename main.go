package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"go-totp-example/http"
	"go-totp-example/services/pquerna"
	"go-totp-example/utils"
)

const appPort = ":3000"

func main() {
	initFlags()

	app := http.NewRouter()

	log.Fatal(app.Listen(appPort))
}

func initFlags() {
	genFlag := flag.Bool("gen", false, "start output current totp codes")
	flag.Parse()

	if *genFlag {
		produceCodes()
	}
}

func produceCodes() {
	srv := pquerna.New()

	go func() {
		for {
			code := srv.GenerateNow(utils.UserSecret)
			fmt.Printf("\n%s - %s", code, time.Now().Format(time.RFC3339))
			time.Sleep(2 * time.Second)
		}
	}()
}
