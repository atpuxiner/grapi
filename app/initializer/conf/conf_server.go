package conf

type ServerConf struct {
	GinMode string `mapstructure:"ginMode"`
	Port    string `mapstructure:"port"`
}
