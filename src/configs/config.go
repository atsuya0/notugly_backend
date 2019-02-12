package config

import (
	"encoding/json"
	"log"
	"os"
)

const message = "Config Error"

type backend struct {
	Port string `json:"port"`
}

type frontend struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type database struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type config struct {
	Backend  backend  `json:"backend"`
	Frontend frontend `json:"frontend"`
	DB       database `json:"database"`
}

var Data config

func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(message, err)
	}

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&Data); err != nil {
		log.Fatalln(message, err)
	}

	if port := os.Getenv("GO_PORT"); port != "" {
		Data.Backend.Port = port
	}
	if user := os.Getenv("MYSQL_USER"); user != "" {
		Data.DB.User = user
	}
	if password := os.Getenv("MYSQL_PASSWORD"); password != "" {
		Data.DB.Password = password
	}
	if host := os.Getenv("MYSQL_HOST"); host != "" {
		Data.DB.Host = host
	}
	if port := os.Getenv("MYSQL_PORT"); port != "" {
		Data.DB.Port = port
	}
}
