package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	instance *Config
)

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Cors     CorsConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type CorsConfig struct {
	AllowOrigins string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
}

type PostgresConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DbName          string
	SslMode         string
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifeTime time.Duration
	MaxConnIdleTime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               int
	Password           string
	Db                 int
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
	IdleCheckFrequency time.Duration
}

type PasswordConfig struct {
	UpperChars    bool
	LowerChars    bool
	InputNum      bool
	InputSpecials bool
	MinLength     int
	MaxLength     int
}

func FindConfigPath(env string) string {
	if env == "docker" {
		return "/app/configs/config-docker"
	} else if env == "production" {
		return "/configs/config-production"
	} else {
		return "../configs/config-development"
	}
}

func LoadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("error occurred while reading config")
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("file does not exist")
		}
		return nil, fmt.Errorf("an error was occurred")
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	config := Config{}
	err := v.Unmarshal(&config)
	if err != nil {
		log.Fatal("an error occurred while parsing config")
	}

	return &config, nil
}

func GetConfig() *Config {
	if instance != nil {
		return instance
	}

	path := FindConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(path, "yml")
	if err != nil {
		log.Fatal("can not load config")
	}

	cnf, err := ParseConfig(v)
	if err != nil {
		log.Fatal("can not parse config")
	}

	instance = cnf

	return instance
}
