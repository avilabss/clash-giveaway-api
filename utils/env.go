package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ClashApiJwt     string `mapstructure:"CLASH_API_JWT"`
	RunUri          string `mapstructure:"RUN_URI"`
	ClashApiBaseUri string `mapstructure:"CLASH_API_BASE_URI"`
}

func LoadEnv(envPath string) (*Env, error) {
	var env Env

	viper.AddConfigPath(envPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = viper.Unmarshal(&env)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &env, nil
}
