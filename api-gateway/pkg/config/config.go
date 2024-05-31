package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"PORT"`
	HireoAuth  string `mapstructure:"Hireo_Auth"`
	HireoJob   string `mapstructure:"Hireo_Job"`
	ChatSvcUrl string `mapstructure:"CHAT_SVC_URL"`

	KafkaPort  string `mapstructure:"KAFKA_PORT"`
	KafkaTopic string `mapstructure:"KAFKA_TOPIC"`
}

var envs = []string{
	"PORT", "Hireo_Admin", "Hireo_Job",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil

}
