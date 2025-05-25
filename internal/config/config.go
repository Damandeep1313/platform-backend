package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Server struct {
        Port int
    }
    // add other config fields here
}

func LoadConfig(path string) (*Config, error) {
    v := viper.New()
    v.SetConfigFile(path)
    err := v.ReadInConfig()
    if err != nil {
        return nil, err
    }
    var cfg Config
    err = v.Unmarshal(&cfg)
    if err != nil {
        return nil, err
    }
    return &cfg, nil
}
