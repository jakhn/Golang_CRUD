package config

type Config struct {
	HTTPPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
}

func Load() Config {

	var cfg Config

	cfg.HTTPPort = ":4000"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "jakh"
	cfg.PostgresDatabase = "my_db"
	cfg.PostgresPassword = "00"
	cfg.PostgresPort = "5432"

	return cfg
}
