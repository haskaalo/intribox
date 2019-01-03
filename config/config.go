package config

import (
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// intribox global configuration
var (
	Debug bool

	Server struct {
		Port        int
		MaxSongSize int64
	}

	Database struct {
		User     string
		Database string
		Password string
		Host     string
		Port     int
		SSLMode  string
	}

	Redis struct {
		Prefix   string
		Host     string
		Port     int
		Password string
		Database int
	}

	Client struct {
		AssetsPath string
	}

	Storage struct {
		UserDataPath string
		RemoteName   string
	}
)

var (
	dir, _ = os.Getwd()
)

func getEnv(key string, fallback string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		val = fallback
	}
	return val
}

func init() {
	// log config
	zerolog.TimeFieldFormat = ""
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cfg, err := ini.Load(getEnv("CONFIG_PATH", "./intribox_config.ini"))
	if err != nil {
		log.Fatal().Err(err).Msg("Config cannot be loaded")
	}

	Debug = cfg.Section("").Key("debug").MustBool(false)
	if Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true})
	}

	Server.Port = cfg.Section("Server").Key("port").MustInt(8080)
	Server.MaxSongSize = cfg.Section("Server").Key("maxsongsize").MustInt64(2000000000)

	Database.User = cfg.Section("Database").Key("user").MustString("postgres")
	Database.Database = cfg.Section("Database").Key("database").MustString("intribox")
	Database.Password = cfg.Section("Database").Key("password").MustString("")
	Database.Host = cfg.Section("Database").Key("host").MustString("localhost")
	Database.Port = cfg.Section("Database").Key("port").MustInt(5432)
	Database.SSLMode = cfg.Section("Database").Key("sslmode").In("disable", []string{"enable", "disable"})

	Redis.Prefix = cfg.Section("Redis").Key("prefix").MustString("intribox")
	Redis.Host = cfg.Section("Redis").Key("host").MustString("localhost")
	Redis.Port = cfg.Section("Redis").Key("port").MustInt(6379)
	Redis.Password = cfg.Section("Redis").Key("password").MustString("")
	Redis.Database = cfg.Section("Redis").Key("Database").MustInt(0)

	Client.AssetsPath = cfg.Section("Client").Key("assetspath").MustString(dir + "/client/dist")

	Storage.UserDataPath = cfg.Section("Storage").Key("userdatapath").MustString(filepath.Join(dir, "/data"))
	Storage.RemoteName = cfg.Section("Storage").Key("remotename").MustString("local")
}
