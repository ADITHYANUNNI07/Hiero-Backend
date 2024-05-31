package config

import "github.com/spf13/viper"

type Config struct {
	DBUri          string `mapstructure:"DB_URI"`
	Port           string `mapstructure:"PORT"`
	Explorite_Auth string `mapstructure:"Explorite_Auth"`
	KafkaBrokers   string `mapstructure:"KAFKA_BROKERS"`
	KafkaTopic     string `mapstructure:"KAFKA_TOPIC"`
}

var envs = []string{
	"DB_URI", "PORT", "KAFKA_BROKERS", "KAFKA_TOPIC", "Explorite_Auth",
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
