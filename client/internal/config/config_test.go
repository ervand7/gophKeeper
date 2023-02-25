package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServerAddressDefaultValue(t *testing.T) {
	assert.NotContainsf(t, os.Args, "-a", "flag -a was set")
	assert.Equal(t, os.Getenv("SERVER_ADDRESS"), "")
	expectedServerAddress := ":8080"
	assert.Equal(t, GetServerAddress(), expectedServerAddress)
}

func TestGetServerAddressFromEnv(t *testing.T) {
	serverAddress := ":5000"
	err := os.Setenv("SERVER_ADDRESS", serverAddress)
	assert.NoError(t, err)
	assert.Equal(t, GetServerAddress(), serverAddress)
}

func TestGetServerAddressFlagPriority(t *testing.T) {
	serverAddressFlag := ":5000"
	serverAddressEnv := ":7777"
	os.Args = []string{"test", "-a", serverAddressFlag}
	err := os.Setenv("SERVER_ADDRESS", serverAddressEnv)
	assert.NoError(t, err)
	assert.Equal(t, GetServerAddress(), serverAddressFlag)
}
