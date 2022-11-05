package main

import (
	"project2-golang/configs"
	"project2-golang/routers"
)

func main() {
	PORT := ":3000"
	db := configs.StartDB()

	router := routers.StartServer(db)

	router.Run(PORT)
}
