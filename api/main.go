package main

import (
	"os"
	"squall/cmd"
)

func main() {
	app := cmd.NewMainApp()
	_ = cmd.RunMainApp(app, os.Args...)
}
