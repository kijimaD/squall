package models

import (
	"squall/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	t.Run("sqlite3/有効なパス", func(t *testing.T) {
		originConfig := config.Config
		config.Config.DBDriver = "sqlite3"
		config.Config.DBConnectionStr = ":memory:"
		config.Config.DBConnectionRetryAttempts = 1
		defer func() {
			config.Config = originConfig
		}()

		db, err := ConnectDB()
		assert.NotNil(t, db)
		assert.NoError(t, err)
	})
	t.Run("対応してないドライバだと、エラーを返す", func(t *testing.T) {
		originConfig := config.Config
		config.Config.DBDriver = "invalid_driver"
		config.Config.DBConnectionStr = ""
		config.Config.DBConnectionRetryAttempts = 1
		defer func() {
			config.Config = originConfig
		}()

		_, err := ConnectDB()
		assert.Error(t, err)
	})
}

func TestGetDB(t *testing.T) {
	originConfig := config.Config
	config.Config.DBDriver = "sqlite3"
	config.Config.DBConnectionStr = ":memory:"
	defer func() {
		config.Config = originConfig
	}()

	db := GetDB()
	assert.NotNil(t, db)
}
