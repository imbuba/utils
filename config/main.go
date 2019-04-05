package config

import (
	"encoding/json"
	"os"
)

// LoadConfiguration универсальный загрузчик конфы
func LoadConfiguration(filename string, config interface{}) (interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)

	if err != nil {
		return config, err
	}

	return config, nil
}
