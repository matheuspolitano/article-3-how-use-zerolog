package config

type Config struct {
	Server ServerConfig
	Logger LoggerConfig
}

type ServerConfig struct {
	AppVersion string
	Port       string
	Mode       string
	Debug      bool
}

type LoggerConfig struct {
	Development        bool
	DisableCaller      bool
	DisableStackTRACER bool
	Level              string
}
