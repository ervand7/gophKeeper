package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigDefaultValues(t *testing.T) {
	assert.NotContainsf(t, os.Args, "-a", "flag -a was set")
	assert.Equal(t, os.Getenv("SERVER_ADDRESS"), "")
	expectedServerAddress := ":8080"
	assert.Equal(t, getConfig().ServerAddress, expectedServerAddress)
}

func TestGetConfigFromEnv(t *testing.T) {
	serverAddress := ":5000"
	err := os.Setenv("SERVER_ADDRESS", serverAddress)
	assert.NoError(t, err)
	assert.Equal(t, getConfig().ServerAddress, serverAddress)
}

func TestGetConfigFlagPriority(t *testing.T) {
	serverAddressFlag := ":5000"
	serverAddressEnv := ":7777"
	os.Args = []string{"test", "-a", serverAddressFlag}
	err := os.Setenv("SERVER_ADDRESS", serverAddressEnv)
	assert.NoError(t, err)
	assert.Equal(t, getConfig().ServerAddress, serverAddressFlag)
}

func TestGetServerAddress(t *testing.T) {
	value := "hello"
	cacheServerAddress = value
	assert.Equal(t, value, cacheServerAddress)
	assert.Equal(t, value, GetServerAddress())
}

func TestGetDatabaseDSN(t *testing.T) {
	value := "hello"
	cacheDatabaseDSN = value
	assert.Equal(t, value, cacheDatabaseDSN)
	assert.Equal(t, value, GetDatabaseDSN())
}
