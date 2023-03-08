package config

import (
	"crypto/sha512"

	"encoding/json"
	"fmt"
	"os"

	"github.com/gabereiser/datalab/log"
	"github.com/google/uuid"
)

type Configuration struct {
	Domain      string  `json:"domain_name"`
	SecretKey   string  `json:"secret_key"`
	Timeout     int     `json:"timeout"`
	ClusterHost *string `json:"cluster_host,omitempty"`

	DatabaseUser     string `json:"database_user"`
	DatabasePassword string `json:"database_password"`
	DatabaseUrl      string `json:"database_url"`
	DatabasePort     int    `json:"database_port"`
	DatabaseName     string `json:"database_name"`
}

var Config *Configuration = LoadConfig()

func LoadConfig() *Configuration {
	data, _ := os.ReadFile("config.json")
	if len(data) == 0 {
		data, _ = os.ReadFile("/etc/datalab/config.json")
		if len(data) == 0 {
			log.Err("unable to find configuration file config.json here or in /etc/datalab")
		}
	}
	var c Configuration
	err := json.Unmarshal(data, &c)
	if err != nil {
		log.Err("%v", err)
	}
	c.Domain = env("DATALAB_DOMAIN", "localhost")
	c.SecretKey = env("DATALAB_SECRET_KEY", genSecretKey(c.SecretKey))
	SaveConfig(&c)
	return &c
}

func SaveConfig(config *Configuration) {
	d, e := json.MarshalIndent(config, "", "	")
	if e != nil {
		log.Err("%v", e)
	}
	e = os.WriteFile("config.json", d, 0644)
	if e != nil {
		log.Err("%v", e)
	}
}

func env(name string, dfault string) string {
	s := os.Getenv(name)
	if s == "" {
		return dfault
	} else {
		return s
	}
}

func genSecretKey(secretKey string) string {
	if secretKey == "" {
		id := uuid.New()
		hash := sha512.New()
		hash.Write([]byte(id.String()))
		s := hash.Sum(nil)
		return fmt.Sprintf("%x", s)
	}
	return secretKey

}
