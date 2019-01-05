package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/haskaalo/intribox/config"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq" // Import lib/pq for database
	"github.com/rs/zerolog/log"
)

var (
	r  *redis.Client
	db *sqlx.DB

	// ErrRecordNotFound Row not found after doing query
	ErrRecordNotFound = errors.New("Record not found in database")

	// ErrRecordAlreadyExist violate UNIQUE constraint or a similar row already exist
	ErrRecordAlreadyExist = errors.New("Record already exist in database")
)

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

func knownDatabaseError(err error) error {
	if err == sql.ErrNoRows {
		return ErrRecordNotFound
	}

	if pgerr, ok := err.(*pq.Error); ok { // https://github.com/lib/pq/blob/master/error.go#L78 and https://www.postgresql.org/docs/9.3/errcodes-appendix.html
		if pgerr.Code.Name() == "unique_violation" {
			return ErrRecordAlreadyExist
		}
	}

	return err
}
