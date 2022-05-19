package dbmodel

import (
	"context"
	"flag"
	"fmt"
	"time"

	applogger "resource-service/utils/logging"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log *zerolog.Logger = applogger.GetInstance()

var cfg Config

type Config struct {
	Port int
	Env  string
	Db   struct {
		DSN string
	}
}

func init() {
	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	// Read the DSN value from the db-dsn command-line flag into the config struct. We
	// default to using our development DSN if no flag is provided.
	flag.StringVar(&cfg.Db.DSN, "db-dsn", "postgres://resource:hiclass@12@localhost/resource", "PostgreSQL DSN")
	flag.Parse()
}

func SetupDB() *gorm.DB {

	db, err := openDB(cfg)
	if err != nil {
		log.Error().Msg(fmt.Sprint(err))
	}
	// Defer a call to db.Close() so that the connection pool is closed before the
	// main() function exits.

	// defer db.Close()

	// Also log a message to say that the connection pool has been successfully
	// established.
	log.Printf("database connection pool established")

	return db

}

// The openDB() function returns a sql.DB connection pool.
func openDB(cfg Config) (*gorm.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config
	// struct.
	db, err := gorm.Open(postgres.Open(cfg.Db.DSN), &gorm.Config{})

	if err != nil {
		log.Panic().Msgf("Error connecting to the database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panic().Msgf("Error getting GORM DB definition")
	}
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(25)
	// Create a context with a 5-second timeout deadline.
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Return the sql.DB connection pool.
	return db, nil
}
