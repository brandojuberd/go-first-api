package models

import (
	"fmt"
	"net/url"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabaseConnection() {
	// err := godotenv.Load("development.env")
	// if(err != nil){
	// 	fmt.Println(err)
	// }
	// user := "postgres"
	// password := "postgres"
	// host := "localhost"
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	fmt.Println(host, port, dbName, user, password)

	fmt.Println("Success connect/create to db: " + dbName)

	dsn := url.URL{
		User:     url.UserPassword(user, password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", host, port),
		Path:     dbName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
