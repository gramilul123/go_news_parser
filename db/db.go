package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"go_news_parser/news"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func init() {
	dbConn := GetDB()
	dbConn.Debug().AutoMigrate(&news.News{})
}

// CreateDbConnect function creates DB connection
func CreateDbConnect(host, user, password, name string) *gorm.DB {
	dbFunc := func() {
		dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, name, password)
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		db = conn
	}
	once.Do(dbFunc)

	return db
}

// GetDB function returns DB connection
func GetDB() *gorm.DB {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	name := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	dbConn := CreateDbConnect(host, user, password, name)

	return dbConn
}
