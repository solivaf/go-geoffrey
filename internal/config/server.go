package config

type ServerConfig struct {
	P string `yaml:"port"`
}

func (sc *ServerConfig) Port() string {
	return sc.P
}
