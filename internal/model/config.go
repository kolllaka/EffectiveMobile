package model

type Config struct {
	IsDebug  bool
	LogLevel string
	Postgres struct {
		Host            string
		Port            string
		User            string
		Password        string
		DBName          string
		SSlmode         string
		MigrationFolder string
	}
	HTTP struct {
		Host string
		Port string
	}
	SongApi struct {
		Scheme    string
		Host      string
		PathsInfo string
	}
	Song struct {
		SongOnPage int
	}
}
