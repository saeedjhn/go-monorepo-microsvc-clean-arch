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

type MYSQL struct {
	Host string `env:"MYSQL_HOST"`
}

type CORS struct {
	AllowOrigins     []string `env:"CORS_ALLOW_ORIGINS"`
	AllowMethods     []string `env:"CORS_ALLOW_METHODS"`
	AllowHeaders     []string `env:"CORS_ALLOW_HEADERS"`
	AllowCredentials bool     `env:"CORS_ALLOW_CREDENTIALS"`
}

type Config struct {
	HTTPServer HTTPServer
	MySQL      MYSQL
}
