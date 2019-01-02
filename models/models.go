package models

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/haskaalo/intribox/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import lib/pq for database
	"github.com/rs/zerolog/log"
)

var (
	r  *redis.Client
	db *sqlx.DB
)

// Querier sql.DB but usable for transactions
type Querier interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
}

func init() {
	initiateDatabase()
	initiateRedis()
}

func initiateDatabase() {
	var err error

	log.Info().Str("database", config.Database.Database).Msg("Connecting to PostgreSQL database")

	connString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%v sslmode=%s connect_timeout=5", config.Database.User, config.Database.Password, config.Database.Database, config.Database.Host, config.Database.Port, config.Database.SSLMode)
	db, err = sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("Failed to connect to PostgreSQL database")
	}

	log.Info().Str("database", config.Database.Database).Msg("Connected to PostgreSQL database!")
}

func initiateRedis() {
	r = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})
}
