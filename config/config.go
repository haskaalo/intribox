package config

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-ini/ini"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// intribox global configuration
var (
	Debug bool

	Server struct {
		Hostname     string // can be set to localhost:port
		Port         int
		MaxMediaSize int64
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

	Aws struct {
		Endpoint  string
		Bucket    string
		Region    string
		AccessKey string
		SecretKey string
	}

	AwsSession *session.Session
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
	Server.Hostname = cfg.Section("Server").Key("hostname").MustString("localhost:" + strconv.Itoa(Server.Port))
	Server.MaxMediaSize = cfg.Section("Server").Key("maxmediasize").MustInt64(2000000000)

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

	Aws.Endpoint = cfg.Section("S3").Key("endpoint").MustString("http://127.0.0.1:9000") // Default value equal to the s3 testing server
	Aws.Bucket = cfg.Section("S3").Key("bucket").MustString("testbucket")                // Default value equal to the s3 testing server
	Aws.Region = cfg.Section("S3").Key("region").MustString("us-east1")
	Aws.AccessKey = cfg.Section("S3").Key("accesskey").MustString("DevAccessKey")
	Aws.SecretKey = cfg.Section("S3").Key("secretkey").MustString("DevSecretKey")
	AwsSession, _ = session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(Aws.AccessKey, Aws.SecretKey, ""),
		Endpoint:    aws.String(Aws.Endpoint),
		Region:      aws.String(Aws.Region),
	})

	if Debug {
		session := s3.New(AwsSession)
		_, _ = session.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(Aws.Bucket),
		})
	}

	if Debug { // When creating a local S3 bucket with localstack, it use localhost so <bucketname>.localhost/<key> doesn't exist, but localhost/<key> does.
		// AwsSession.Config = AwsSession.Config.WithS3ForcePathStyle(true)
	}
}
