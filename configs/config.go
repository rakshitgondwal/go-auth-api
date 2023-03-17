package configs


type ConfigDatabase struct {
    Port     string `yaml:"port" env:"PORT" env-default:"1323"`
    Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
}

var Cfg ConfigDatabase

