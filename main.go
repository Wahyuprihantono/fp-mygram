package main

import (
	"os"
	"project2-golang/configs"
	"project2-golang/routers"
)

func main() {
	port := os.Getenv("DB_PORT")
	db := configs.StartDB()

	router := routers.StartServer(db)

	router.Run(":" + port)
}
