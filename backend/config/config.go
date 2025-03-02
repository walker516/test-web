package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DatabaseConfig struct {
	DBName    string `json:"db_name"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Charset   string `json:"charset"`
	ParseTime bool   `json:"parse_time"`
	Loc       string `json:"loc"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Config struct {
	DatabaseConfigs map[string]DatabaseConfig `json:"database"`
	ServerConfig    ServerConfig              `json:"server"`
}

var ConfigData Config

func InitConfig() error {
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		configFilePath = "config.dev.json"
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	if err = json.Unmarshal(data, &ConfigData); err != nil {
		return fmt.Errorf("error unmarshaling config data: %w", err)
	}

	// 環境変数があれば上書き
	if port := os.Getenv("SERVER_PORT"); port != "" {
		ConfigData.ServerConfig.Port = port
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		dbConfig := ConfigData.DatabaseConfigs["my_schema"]
		dbConfig.User = dbUser
		ConfigData.DatabaseConfigs["my_schema"] = dbConfig
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		dbConfig := ConfigData.DatabaseConfigs["my_schema"]
		dbConfig.Password = dbPassword
		ConfigData.DatabaseConfigs["my_schema"] = dbConfig
	}

	return nil
}
