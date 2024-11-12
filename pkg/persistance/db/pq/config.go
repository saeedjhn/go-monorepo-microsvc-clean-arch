package pq

import "time"

type Config struct {
	Host            string        `env:"host"`
	Port            string        `env:"port"`
	Username        string        `env:"username"`
	Password        string        `env:"password"`
	Database        string        `env:"database"`
	SSLMode         string        `env:"ssl_mode"`
	MaxIdleConns    int           `env:"max_idle_conns"`
	MaxOpenConns    int           `env:"max_open_conns"`
	ConnMaxLiftTime time.Duration `env:"conn_max_life_time"`
}
