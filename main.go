package main

import (
	"log"

	"github.com/pradiptarana/book-online-store/app"
)

func main() {
	if err := app.SetupServer().Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
