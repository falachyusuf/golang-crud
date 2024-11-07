package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

var (
	ServerPort string
	ServerTimeout int
)

func init() {
	ServerPort = doGetEnv("SERVER_PORT", "8080")
	ServerTimeout = doGetEnvInt("SERVER_TIMEOUT", 5000)
}

// function to retrieve env variable of server port values
func doGetEnv(key string, def string) string {
	// Check if the value of the key is ""/empty
	if val := os.Getenv(key); val != "" {
		return val
	}
	// Else return default
	return def
}

// function to retrieve env variable of server timeout port values
func doGetEnvInt(key string, def int) int {
	// Check if the value of the key is ""/empty
	if val := os.Getenv(key); val != "" {
		/* Convert the value to integer and return it if successful,
		otherwise return the default value */
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return def
		}
		return intVal
	}
	// Else return default
	return def
}

// function to get the server port to running the server
func DoGetServerPort() string {
	return ServerPort
}

// function to get server timeout for the server
func DoGetServerTimeout() int {
	return ServerTimeout
}