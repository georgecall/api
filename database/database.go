package database

import (
	"time"

	"github.com/georgecall/api/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
        "github.com/theGOURL/warning"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
    warning.PRINT_DEFAULT_ERRORS(err,"Could not connect to the Postgres Database")

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func CloseConn() error {
	config, err := db.DB()
    warning.FATAL_ERROR(err)

	err = config.Close()
    warning.FATAL_ERROR(err)

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
