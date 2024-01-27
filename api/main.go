package main

import (
	"log"
	"os"
	"squall/cmd"
)

func main() {
	app := cmd.NewMainApp()
	err := cmd.RunMainApp(app, os.Args...)
	log.Fatal(err)
}
