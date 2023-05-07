package config

type System struct {
	Env  string `yaml:"env"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
