package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PostgreSQLConfig struct {
	Username string `env:"POSTGRES_USERNAME,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	Host     string `env:"POSTGRES_HOST,required"`
	Port     string `env:"POSTGRES_PORT,required"`
	DBName   string `env:"POSTGRES_DBNAME,required"`
}
type AppConfig struct {
	ServerPort string `env:"PORT,required"`
	JWTSecret  string `env:"JWT_SECRET,required"`
	Postgres   PostgreSQLConfig
}

var conf *AppConfig

func GetAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pg := PostgreSQLConfig{
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DBNAME"),
	}

	conf = &AppConfig{
		ServerPort: os.Getenv("PORT"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		Postgres:   pg,
	}
	return conf
}
