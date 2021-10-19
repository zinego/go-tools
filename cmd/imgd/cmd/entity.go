package cmd

var imgdCfg imgdConfiguration

type imgdConfiguration struct {
	Auth           authentication `mapstructure:"auth"`
	ImgRespository string         `mapstructure:"img-respository"`
	ImgUrlPrefix   string         `mapstructure:"img-url-prefix"`
	RemoteName     string         `mapstructure:"remote-name"`
}

type authentication struct {
	Method         string `mapstructure:"method"`
	Password       string `mapstructure:"password"`
	Username       string `mapstructure:"username"`
	PrivateKeyPath string `mapstructure:"private-key-path"`
}
