package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

func Dd(data any)  {
	
	fmt.Println(data);
	os.Exit(0);
}


func GetConfig(key string) (interface{}, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	value, exists := config[key]
	if !exists {
		return nil, fmt.Errorf("key %s not found in config", key)
	}

	return value, nil
}