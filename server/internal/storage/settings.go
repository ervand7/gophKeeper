package storage

import (
	"database/sql"
	"path/filepath"
	"runtime"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"

	"gophkeeper/server/internal/config"
	"gophkeeper/server/internal/logger"
)

const (
	maxOpenConnections    = 20
	maxIdleConnections    = 20
	connMaxIdleTimeSecond = 30
	connMaxLifetimeSecond = 2
)

// Database for manage DB settings.
type Database struct {
	conn *sql.DB
}

// Run launches DB.
func (db *Database) Run() {
	err := db.ConnStart()
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	db.setConnPool()
	err = db.migrate()
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
}

// ConnStart opens DB connection.
func (db *Database) ConnStart() (err error) {
	conn, err := goose.OpenDBWithDriver("pgx", config.GetDatabaseDSN())
	if err != nil {
		return err
	}
	db.conn = conn
	return nil
}

// ConnClose closes DB connection.
func (db *Database) ConnClose() (err error) {
	err = db.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

// Ping is used for check if DB connection is alive.
func (db *Database) Ping() (err error) {
	err = db.conn.Ping()
	if err != nil {
		return err
	}
	return nil
}

// closeRows closes db.conn.Query rows.
func (db *Database) closeRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		logger.Logger.Error(err.Error())
	}
}

func (db *Database) setConnPool() {
	db.conn.SetMaxOpenConns(maxOpenConnections)
	db.conn.SetMaxIdleConns(maxIdleConnections)
	db.conn.SetConnMaxIdleTime(time.Second * connMaxIdleTimeSecond)
	db.conn.SetConnMaxLifetime(time.Minute * connMaxLifetimeSecond)
}

func (db *Database) migrate() (err error) {
	if err = goose.Run("up", db.conn, getMigrationsDir()); err != nil {
		return err
	}
	return nil
}

func getMigrationsDir() string {
	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)
	migrationsDir := filepath.Join(currentDir, "/../../../migrations")
	return migrationsDir
}
