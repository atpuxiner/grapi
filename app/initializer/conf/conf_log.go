package conf

type LogConf struct {
	Debug     bool `mapstructure:"debug"`
	MaxSize   int  `mapstructure:"maxSize"`
	MaxAge    int  `mapstructure:"maxAge"`
	MaxBackup int  `mapstructure:"maxBackup"`
	LocalTime bool `mapstructure:"localTime"`
	Compress  bool `mapstructure:"compress"`
}
