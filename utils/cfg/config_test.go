package cfg

import (
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestReadConfiguration(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		var imgdCfg ImgdConfiguration
		var err = ReadConfiguration("imgd", &imgdCfg)
		assert.Equal(t, nil, err)
		pretty.Println(imgdCfg)
	})
	t.Run("error", func(t *testing.T) {
		var imgdCfg ImgdConfiguration
		var err = ReadConfiguration("imgd.ya", &imgdCfg)
		assert.NotEqual(t, `nil`, err)
		pretty.Println(imgdCfg)
	})
}

type ImgdConfiguration struct {
	Auth           authentication `mapstructure:"auth"`
	ImgRespository string         `mapstructure:"img_respository"`
}

type authentication struct {
	Method         string `mapstructure:"method"`
	Password       string `mapstructure:"password"`
	Username       string `mapstructure:"username"`
	PrivateKeyPath string `mapstructure:"private_key_path"`
}
