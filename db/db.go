package db

import (
	"fmt"
	"log"
	"time"

	"mobile-rest-api/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var dbSafe *gorm.DB
var err error

//Init - Database init
func Init() {

	config.Init()
	var cfg = config.Get()

	user := cfg.Database.User
	password := cfg.Database.Pass
	host := cfg.Database.Host
	port := cfg.Database.Port
	database := cfg.Database.DBName
	appName := cfg.AppName

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable application_name=%s",
		user,
		password,
		host,
		port,
		database,
		appName,
	)

	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	db.DB().SetConnMaxLifetime(time.Minute * 5)
}

//Init - Database init
func InitSafeCityDB() {

	config.Init()
	var cfg = config.Get()

	user := cfg.DatabaseSafeCity.User
	password := cfg.DatabaseSafeCity.Pass
	host := cfg.DatabaseSafeCity.Host
	port := cfg.DatabaseSafeCity.Port
	database := cfg.DatabaseSafeCity.DBName
	appName := cfg.AppName

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable application_name=%s",
		user,
		password,
		host,
		port,
		database,
		appName,
	)

	dbSafe, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	dbSafe.DB().SetConnMaxLifetime(time.Minute * 5)
}

//GetDB - get DB
func GetSafeCityDB() *gorm.DB {
	return dbSafe
}

//CloseDB - close DB
func CloseSafeCityDB() {
	dbSafe.Close()
}

//GetDB - get DB
func GetDB() *gorm.DB {
	return db
}

//CloseDB - close DB
func CloseDB() {
	db.Close()
}
