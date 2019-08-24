package config

type AppConfig struct {
	Gc *GitConfig    `yaml:"git"`
	Sc *ServerConfig `yaml:"server"`
}

func (c *AppConfig) GitConfig() *GitConfig {
	return c.Gc
}

func (c *AppConfig) ServerConfig() *ServerConfig {
	return c.Sc
}
