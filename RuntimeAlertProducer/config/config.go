package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

type Config struct {
	KafkaConsumer KafkaConsumerConfig
	KafkaProducer KafkaProducerConfig
}

type KafkaConsumerConfig struct {
	Broker           string
	Protocol         string
	Mechanisms       string
	SaslUsername     string
	SaslPassword     string
	AutoOffsetReset  string
	ConsumerGroup    string
	SessionTimeout   int
	Topics           []string
	FetchMinBytes    int
	PollTimeout      int
	EnableAutoCommit bool
	NoOfRetry        int
}

type KafkaProducerConfig struct {
	Brokers    string
	Protocol   string
	Mechanisms string
	UserName   string
	Password   string
	BatchSize  int
	LingerMs   int
}

var once sync.Once

var c *Config
var configErr error

func Get() (*Config, error) {
	once.Do(func() {
		c, configErr = loadConfig()
	})
	return c, configErr

}

func loadConfig() (*Config, error) {
	configPath := getConfigPath(os.Getenv("config"))

	v := viper.New()
	v.SetConfigName(configPath)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	//setEnvSecrets(v)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	var c *Config
	err := v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func getConfigPath(env string) string {
	log.Printf("Loading config file for %s environment", env)
	if env == "local" {
		return "../config/config-local"
	}
	return "../config/config"
}
