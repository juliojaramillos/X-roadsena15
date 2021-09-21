package main

import (
	"unal.edu.co/rest-example/app"
	"unal.edu.co/rest-example/config"
	"unal.edu.co/rest-example/database"
)

func main() {
	config.InitConfiguration()
	database.InitDatabase()
	app.StartApp()
}
