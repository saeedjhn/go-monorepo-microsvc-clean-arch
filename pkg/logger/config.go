package logger

type Config struct {
	Filename   string `env:"filename"`
	MaxSize    int    `env:"max_size"`
	MaxBackups int    `env:"max_backups"`
	MaxAge     int    `env:"max_age"`
	LocalTime  bool   `env:"local_time"`
	Compress   bool   `env:"compress"`
}

/* for example:
Filename: "./logs/log.json",
MaxSize:  10, // megabytes
MaxBackups: 10, // megabytes
MaxAge:    30, // days
LocalTime: false,
Compress:  false,
*/
