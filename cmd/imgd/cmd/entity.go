package cmd

var imgdCfg imgdConfiguration

type imgdConfiguration struct {
	Auth           authentication `mapstructure:"auth"`
	ImgRespository string         `mapstructure:"img_respository"`
}

type authentication struct {
	Method         string `mapstructure:"method"`
	Password       string `mapstructure:"password"`
	Username       string `mapstructure:"username"`
	PrivateKeyPath string `mapstructure:"private_key_path"`
}
