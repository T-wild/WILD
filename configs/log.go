package configs

type LogConfig struct {
	Level      string `mapstructure:"Level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackUps int    `mapstructure:"max_backups"`
	Mode       string `mapstructure:"log_mode"`
}
