package configs

type ServerConfig struct {
	Name             string `mapstructure:"name"`
	Mode             string `mapstructure:"mode"`
	Version          string `mapstructure:"version"`
	Port             int    `mapstructure:"port"`
	*AuthConfig      `mapstructure:"auth"`
	*LogConfig       `mapstructure:"log"`
	*MySQLConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
	*SnowFlakeConfig `mapstructure:"snowflake"`
}
