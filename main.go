package main

import (
	"mygram/database"
	"mygram/routers"
)

func init() {

}

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
