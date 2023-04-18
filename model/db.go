package model

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	err := godotenv.Load("config/.env")

	var (
		DBCONN = os.Getenv("POSTGRES_URL")
		DBNAME = os.Getenv("POSTGRES_DB")
		DBPORT = os.Getenv("POSTGRES_PORT")
		DBUSER = os.Getenv("POSTGRES_USER")
		DBPASS = os.Getenv("POSTGRES_PASSWORD")
	)
	fmt.Printf("Connecting to PostgresQL %s %s", DBCONN, DBNAME)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Manila",
		DBCONN,
		DBUSER,
		DBPASS,
		DBNAME,
		DBPORT,
	)
	fmt.Printf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Manila",
		DBCONN,
		DBUSER,
		DBPASS,
		DBNAME,
		DBPORT)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			},
		),
	})

	if err != nil {
		fmt.Printf("Error connecting to the DB: %s\n", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Error could not ping database: %s\n", err.Error())
	}
	sqlDB.Ping()

	// psqlDB.SetMaxIdleConns(3)
	// psqlDB.SetMaxOpenConns(1)
	// psqlDB.SetMaxConnMaxLifetime(time.Hour)
}
