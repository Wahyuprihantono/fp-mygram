package configs

import (
	"fmt"
	"log"
	"os"
	"project2-golang/helpers"
	"project2-golang/models"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB() *gorm.DB {
	if os.Getenv("APP_NEW") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("PORT")

	dataSourceName := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	helpers.PanicIfError(err)

	err = autoMigrate(db)
	helpers.PanicIfError(err)

	log.Default().Println("connection db succcess")

	return db
}

func autoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}, &models.SocialMedia{}, &models.Photo{}, &models.Comment{}); err != nil {
		return err
	}

	return nil
}
