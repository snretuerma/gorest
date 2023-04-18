package model

import (
	"fmt"
	"log"
	"os"
	"time"

	zlog "github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	var (
		DBCONN = os.Getenv("POSTGRES_URL")
		DBNAME = os.Getenv("POSTGRES_DB")
		DBPORT = os.Getenv("POSTGRES_PORT")
		DBUSER = os.Getenv("POSTGRES_USER")
		DBPASS = os.Getenv("POSTGRES_PASSWORD")
	)

	zlog.Info().Msgf("Connecting to PostgresQL %s %s", DBCONN, DBNAME)
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		DBUSER,
		DBPASS,
		DBCONN,
		DBPORT,
		DBNAME,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			},
		),
	})

	if err != nil {
		zlog.Fatal().Msgf("Failed to connect to database: %v", err)
	}

	psqlDB, err := db.DB()
	if err != nil {
		zlog.Fatal.Msgf("Unable to ping database %s %s", DBCONN, DBNAME)
	}

	psqlDB.SetMaxIdleConns(3)
	psqlDB.SetMaxOpenConns(1)
	psqlDB.SetMaxConnMaxLifetime(time.Hour)
}
