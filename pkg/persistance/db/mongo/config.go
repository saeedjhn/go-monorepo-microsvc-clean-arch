package mongo

type Config struct {
	Host     string `env:"host"`
	Port     string `env:"port"`
	Username string `env:"username"`
	Password string `env:"password"`
	Database string `env:"database"`
}
