package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// stroing db
var DB *gorm.DB

// storing connection
var dbMap map[string]func()

// initializing all database connection
func init() {
	fmt.Println("Start Initializing Database Connection")
	dbMap = make(map[string]func())

	dbMap["mysql"] = MySQLConnect
	dbMap["postgres"] = PostGresConnect

	fmt.Println("Initializing Database Connection Complete")

}

// connection db according connection
func GetDB() *gorm.DB {
	godotenv.Load()
	connection := os.Getenv("DB_CONNECTION")
	if connection == "" {
		connection = "mysql"
	}
	connect, exist := dbMap[connection]
	if exist {
		connect()
	} else {
		panic("No Database Connection Found")
	}
	return DB
}
