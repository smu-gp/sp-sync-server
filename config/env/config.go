package env

import "github.com/spf13/viper"

func NewViperConfig() Config {
	v := &viperConfig{}
	v.Init()
	return v
}

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	Init()
}

type viperConfig struct {
}

func (config *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (config *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (config *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (config *viperConfig) Init() {
	viper.AutomaticEnv()
	viper.SetConfigType(`json`)
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
