package config

import (
	"encoding/json"
	"os"
)

// LoadConfiguration универсальный загрузчик конфы
func LoadConfiguration(filename string, config interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(config); err != nil {
		return err
	}
	return nil
}
