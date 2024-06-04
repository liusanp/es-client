package models

type AppConfig struct {
	Port int `json:"port" yaml:"port"`
}

type ESConfig struct {
	Name     string `json:"name" yaml:"name"`
	Version  string `json:"version" yaml:"version"`
	Address  string `json:"address" yaml:"address"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Selected bool   `json:"selected" yaml:"selected"`
}

type Config struct {
	App AppConfig `json:"app" yaml:"app"`
	ES  struct {
		Conf []ESConfig `json:"conf" yaml:"conf"`
	} `json:"es" yaml:"es"`
}
