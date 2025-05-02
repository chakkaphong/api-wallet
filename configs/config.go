package configs

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Configs struct {
	App     App     `yaml:"setting"`
	Server  Server  `yaml:"server"`
	Service Service `yaml:"service"`
}

type App struct {
	Name string `yaml:"name"`
	Env  string `yaml:"env"`
	Port int    `yaml:"port"`
}

type Server struct {
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Service struct {
	Debug bool `mapstructure:"debug"`
}

func GetConfig() (config Configs) {
	// viper
	viper.AddConfigPath(strings.Join([]string{"./configs"}, "/"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.SetEnvPrefix("TMDS")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %v", err))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}

	config.Server.IdleTimeout = parseDurationWithDefault(viper.GetString("server.idle_timeout"), 1*time.Minute)
	config.Server.ReadTimeout = parseDurationWithDefault(viper.GetString("server.read_timeout"), 10*time.Second)
	config.Server.WriteTimeout = parseDurationWithDefault(viper.GetString("server.write_timeout"), 30*time.Second)

	return config
}

func parseDurationWithDefault(duration string, defaultDuration time.Duration) time.Duration {
	if duration == "" {
		return defaultDuration
	}
	d, err := time.ParseDuration(duration)
	if err != nil {
		return defaultDuration
	}
	return d
}
