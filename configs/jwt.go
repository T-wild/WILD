package configs

type AuthConfig struct {
	JwtSecret string `mapstructure:"jwt_secret"`
	JwtExpire int64  `mapstructure:"jwt_expire"`
}
