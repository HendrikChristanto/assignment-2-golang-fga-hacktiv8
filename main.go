package main

import (
	"assignment-2-golang-hacktiv8/controllers"
	"assignment-2-golang-hacktiv8/config"
	"assignment-2-golang-hacktiv8/routers"
)

var PORT = ":8080"

func main() {
	db := config.StartDB()
	controller := controllers.New(db)

	routers.StartServer(controller).Run(PORT)
}