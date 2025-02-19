package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 0=sp_gateway,
// 1=sp_addrstore,
// 2=sp_addrstore_archive,
var Db []*gorm.DB

// SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() {
	// Use the function directly instead of setting env variables
	dbUser := os.Getenv("PLE_DB_USER")
	dbPass := os.Getenv("PLE_DB_PASS")
	dbHost := os.Getenv("PLE_DB_HOST")
	dbName := os.Getenv("PLE_DB_NAME")

	fmt.Println(dbUser, dbPass, dbHost, dbName)
	// handle global users, use loc=UTC
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		dbUser,
		dbPass,
		dbHost,
		dbName)

	cDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("Error connecting to database" + os.Getenv("ENV") + "dbname" + dbName)
	}

	Db = append(Db, cDb)

}
