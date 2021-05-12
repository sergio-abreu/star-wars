package postgres

type Config struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Database string `env:"DB_DATABASE"`
}
