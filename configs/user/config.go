package user

import "time"

type EnvMode string

func (e EnvMode) String() string {
	return string(e)
}

const (
	Development EnvMode = "development"
	Production  EnvMode = "production"
)

type Application struct {
	Env             EnvMode `env:"ENV_MODE"`
	Debug           bool    `env:"DEBUG"`
	EntropyPassword float64 `env:"ENTROPY_PASSWORD"`
}

type HTTPServer struct {
	Port                    string        `env:"HTTP_PORT"`
	GracefulShutdownTimeout time.Duration `env:"HTTP_GRACEFUL_SHUTDOWN_TIMEOUT"`
}

type GRPCServer struct {
	Network string        `env:"GRPC_NETWORK"`
	Port    string        `env:"GRPC_PORT"`
	Timeout time.Duration `env:"GRPC_TIMEOUT"`
}

type MySQL struct {
	Host            string        `env:"MYSQL_HOST"`
	Port            string        `env:"MYSQL_PORT"`
	Username        string        `env:"MYSQL_USER"`
	Password        string        `env:"MYSQL_PASSWORD"`
	DB              string        `env:"MYSQL_DB"`
	SSLMode         string        `env:"MYSQL_SSL_MODE"`
	MaxIdleConns    int           `env:"MYSQL_MAX_IDLE_CONNS"`
	MaxOpenConns    int           `env:"MYSQL_MAX_OPEN_CONNS"`
	ConnMaxLiftTime time.Duration `env:"MYSQL_MAX_LIFE_TIME"`
}

type Postgres struct {
	Host            string        `env:"POSTGRES_HOST"`
	Port            string        `env:"POSTGRES_PORT"`
	User            string        `env:"POSTGRES_USER"`
	Password        string        `env:"POSTGRES_PASSWORD"`
	DB              string        `env:"POSTGRES_DB"`
	SSLMode         string        `env:"POSTGRES_SSL_MODE"`
	MaxIdleConns    int           `env:"POSTGRES_MAX_IDLE_CONNS"`
	MaxOpenConns    int           `env:"POSTGRES_MAX_OPEN_CONNS"`
	ConnMaxLifetime time.Duration `env:"POSTGRES_CONN_MAX_LIFE_TIME"`
	LogLevel        int           `env:"POSTGRES_LOG_LEVEL"`
}

type Redis struct {
	Host               string        `env:"REDIS_HOST"`
	Port               string        `env:"REDIS_PORT"`
	Password           string        `env:"REDIS_PASSWORD"`
	DB                 int           `env:"REDIS_DB"`
	DialTimeout        time.Duration `env:"REDIS_DIAL_TIMEOUT"`
	ReadTimeout        time.Duration `env:"REDIS_READ_TIMEOUT"`
	WriteTimeout       time.Duration `env:"REDIS_WRITE_TIMEOUT"`
	PoolSize           int           `env:"REDIS_POOL_SIZE"`
	PoolTimeout        time.Duration `env:"REDIS_POOL_TIMEOUT"`
	IdleCheckFrequency time.Duration `env:"REDIS_IDLE_CHECK_FREQUENCY"`
}

type CORS struct {
	AllowOrigins     []string `env:"CORS_ALLOW_ORIGINS"`
	AllowMethods     []string `env:"CORS_ALLOW_METHODS"`
	AllowHeaders     []string `env:"CORS_ALLOW_HEADERS"`
	AllowCredentials bool     `env:"CORS_ALLOW_CREDENTIALS"`
}

type Config struct {
	Application Application
	HTTPServer  HTTPServer
	GRPCServer  GRPCServer
	MySQL       MySQL
	Postgres    Postgres
	Redis       Redis
}
