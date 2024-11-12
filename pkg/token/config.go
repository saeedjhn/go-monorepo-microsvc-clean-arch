package token

import "time"

type Config struct {
	AccessTokenSecret      string        `env:"secret"`
	RefreshTokenSecret     string        `env:"refresh_secret"`
	AccessTokenExpiryTime  time.Duration `env:"access_token_expire_duration"`
	RefreshTokenExpiryTime time.Duration `env:"refresh_token_expire_duration"`
}
