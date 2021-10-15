package cfg

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

func ReadConfiguration(fname string, cfg interface{}) error {
	homeDir, _ := os.UserHomeDir()
	cfgDir, _ := os.UserConfigDir()
	ext := path.Ext(fname)
	fname = strings.TrimSuffix(fname, ext)
	fname = strings.TrimSuffix(fname, ".")
	viper.SetConfigType(ext)
	viper.SetConfigName(fname)
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath(homeDir)
	viper.AddConfigPath(cfgDir)
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error config file: %v", err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return fmt.Errorf("unmarshal conf file: %v", err)
	}
	return nil
}
