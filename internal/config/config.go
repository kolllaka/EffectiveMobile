package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kolllaka/EffectiveMobile/internal/model"
	"github.com/kolllaka/EffectiveMobile/pkg/logging"
)

func GetConfig(logger logging.Logger) *model.Config {
	if err := godotenv.Load(); err != nil {
		logger.Fatalln("No .env file found")
	}

	songOnPage, _ := strconv.Atoi(getEnv("LIMIT_SONG_ON_PAGE", "10"))

	return &model.Config{
		IsDebug:  getEnv("APP_IS_DEBUG", "false") == "true",
		LogLevel: getEnv("SONG_API_LOG_LEVEL", "info"),
		Postgres: struct {
			Host            string
			Port            string
			User            string
			Password        string
			DBName          string
			SSlmode         string
			MigrationFolder string
		}{
			Host:            getEnv("PSQL_HOST", ""),
			Port:            getEnv("PSQL_PORT", "5432"),
			User:            getEnv("PSQL_USERNAME", ""),
			Password:        getEnv("PSQL_PASSWORD", ""),
			DBName:          getEnv("PSQL_DATABASE", ""),
			SSlmode:         getEnv("PSQL_SSLMODE", "disable"),
			MigrationFolder: getEnv("PSQL_MIGRATION_FOLDER", "./migrations"),
		},
		HTTP: struct {
			Host string
			Port string
		}{
			Host: getEnv("HTTP_HOST", "localhost"),
			Port: getEnv("HTTP_PORT", "8080"),
		},
		SongApi: struct {
			Scheme    string
			Host      string
			PathsInfo string
		}{
			Scheme:    getEnv("SONG_API_SCHEME", "https"),
			Host:      getEnv("SONG_API_HOST", "api.example.com"),
			PathsInfo: getEnv("SONG_API_PATHS_INFO", "/info"),
		},
		Song: struct {
			SongOnPage int
		}{
			SongOnPage: songOnPage,
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
