package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

type ConfigVar struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type ConfigFile struct {
	Environments []ConfigVar `yaml:"env"`
	Secrets      []ConfigVar `yaml:"secrets"`
}

type ConfigKey string

type Config struct {
	values map[ConfigKey]string
}

var (
	SqliteConnection ConfigKey = "sqlite"
	Port             ConfigKey = "port"
)

func (c *Config) SetData(data []ConfigVar) {
	if len(c.values) == 0 {
		c.values = make(map[ConfigKey]string, 0)
	}

	for _, val := range data {
		c.values[ConfigKey(strings.ToLower(val.Name))] = val.Value
	}
}

func (c Config) GetValue(key ConfigKey) string {
	if val := os.Getenv(strings.ToUpper(string(key))); val != "" {
		log.Println("env: ", val)
		return val
	}

	val, ok := c.values[key]
	if !ok {
		return ""
	}
	return val
}

func uploadConfigFromYaml() ConfigFile {
	valuesFile := os.Getenv("VALUES")
	if valuesFile == "" {
		valuesFile = "./.o3/k8s/values_local.yaml"
	}

	data, err := os.ReadFile(valuesFile)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var config ConfigFile
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return config
}

// NewConfig .
func NewConfig() Config {
	confFile := uploadConfigFromYaml()

	config := Config{}
	config.SetData(confFile.Secrets)
	config.SetData(confFile.Environments)

	return config
}
