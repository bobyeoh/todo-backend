package main

import (
	"log"
	"todo/app"
	"todo/app/routes"
	_ "todo/docs"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var instance *app.Server

func main() {
	instance = app.NewServer()
	routes.Routes(instance)
	err := instance.Start("8888")
	if err != nil {
		log.Fatal("Port is already occupied.")
	}
}
