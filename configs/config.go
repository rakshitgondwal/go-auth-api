package configs

type ConfigDatabase struct {
	Port       string `yaml:"port" env:"PORT" env-default:"1323"`
	Host       string `yaml:"host" env:"HOST" env-default:"localhost"`
	DbUsername string `yaml:"host" env:"DBUSERNAME" env-default:"rakshitgondwal"`
	DbPassword string `yaml:"host" env:"DBPASSWORD" env-default:"rakshitgondwal"`
	JwtSecret  string `yaml:"host" env:"JWTSECRET"  env-default:"rakshitgondwal"`
}

var Cfg ConfigDatabase
