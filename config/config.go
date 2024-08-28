package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	portKey     = "PORT"
	defaultPort = 8080
)

func Port() int {
	num, err := getInt(portKey)
	if err != nil {
		return defaultPort
	}
	return num
}

func getInt(key string) (int, error) {

	v := os.Getenv(key)

	if v == "" {
		return 0, fmt.Errorf("%s is not set", key)
	}

	num, err := strconv.Atoi(v)

	if err != nil {
		return 0, fmt.Errorf("%s is not a number: %s", key, v)
	}

	return num, nil

}

func getString(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", fmt.Errorf("%s is not set", key)
	}
	return v, nil
}
