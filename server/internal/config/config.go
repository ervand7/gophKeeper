// Package config - server configurations.
package config

import (
	"flag"

	"github.com/caarlos0/env/v6"

	"gophkeeper/server/internal/logger"
)

var (
	servAddrFlag    *string
	databaseDSNFlag *string
)

var (
	cacheServerAddress string
	cacheDatabaseDSN   string
)

func init() {
	servAddrFlag = flag.String("a", "", "Server address")
	databaseDSNFlag = flag.String("d", "", "Database source name")
}

type config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080" json:"server_address"`
	DatabaseDSN   string `env:"DATABASE_DSN" json:"database_dsn"`
}

func getConfig() config {
	var cfg config
	err := env.Parse(&cfg)
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}

	flag.Parse()
	if *servAddrFlag != "" {
		cfg.ServerAddress = *servAddrFlag
	}
	if *databaseDSNFlag != "" {
		cfg.DatabaseDSN = *databaseDSNFlag
	}

	return cfg
}

// GetServerAddress gets serverAddress by cache.
func GetServerAddress() string {
	if cacheServerAddress != "" {
		return cacheServerAddress
	}
	cacheServerAddress = getConfig().ServerAddress
	return cacheServerAddress
}

// GetDatabaseDSN gets databaseDSN by cache.
func GetDatabaseDSN() string {
	if cacheDatabaseDSN != "" {
		return cacheDatabaseDSN
	}
	cacheDatabaseDSN = getConfig().DatabaseDSN
	return cacheDatabaseDSN
}
