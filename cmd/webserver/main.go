package main

import (
	"log"
)

func main() {
	app, err := initApp()
	if err != nil {
		log.Fatal(err)
	}

	app.start()
}
