package main

import (
	"fmt"
	"log"

	"github.com/codernishchay/productapi/app"
	"github.com/codernishchay/productapi/config"
)

func main() {
	fmt.Println("Starting the project")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := config.NewConfig()
	app.ConfigAndRunApp(config)

}
