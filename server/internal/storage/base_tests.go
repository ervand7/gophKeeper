package storage

import (
	"github.com/pressly/goose/v3"

	"gophkeeper/server/internal/logger"
)

func Downgrade() {
	db := Database{}
	db.Run()
	defer func() {
		if err := goose.Run("down", db.conn, getMigrationsDir()); err != nil {
			logger.Logger.Error(err.Error())
		}
	}()
}
