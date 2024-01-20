package models

import (
	"fmt"
	"log"
	"os"
	"squall/config"
	"squall/loghelper"
	"sync"

	gormLogger "gorm.io/gorm/logger"

	"github.com/avast/retry-go"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	singletonDB         *gorm.DB
	singletonOnce       sync.Once
	ErrNotFoundDBDriver = fmt.Errorf("対応してないDB Driver: %s", config.Config.DBDriver)
)

type DBDriverType string

const (
	DBDriverSqlite3 = DBDriverType("sqlite3")
)

// ConnectDB は既存のDBに接続する.
func ConnectDB() (*gorm.DB, error) {
	gormConfig := gorm.Config{}
	if config.Config.SilentMode {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	} else {
		gormConfig.Logger = &loghelper.GormSlogLogger{LogLevel: loghelper.GetGormLogLevel(config.Config.LogLevel)}
	}

	var gmdb *gorm.DB
	err := retry.Do(
		func() error {
			var err error
			switch DBDriverType(config.Config.DBDriver) {
			case DBDriverSqlite3:
				gmdb, err = gorm.Open(sqlite.Open(config.Config.DBConnectionStr), &gormConfig)
				// sqliteはデフォルトで外部キー制約を無視してしまうため必要
				if res := gmdb.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
					return res.Error
				}
			default:
				err = ErrNotFoundDBDriver
			}

			return err
		},
		retry.Attempts(config.Config.DBConnectionRetryAttempts),
		retry.Delay(config.Config.DBConnectionRetryDelay),
	)
	if err != nil {
		return nil, fmt.Errorf("データベースとの接続に失敗した: %w", err)
	}

	return gmdb, nil
}

// GetDB DBのsingletonインスタンスを取得する.
func GetDB() *gorm.DB {
	singletonOnce.Do(func() {
		db, err := ConnectDB()
		if err != nil {
			log.Fatal("DBとの接続に失敗した")
		}
		err = db.AutoMigrate(All...)
		if err != nil {
			log.Fatal("マイグレーションが失敗した")
		}
		singletonDB = db
	})

	return singletonDB
}

// NewSQLiteDB は新規DBを作成して接続する.
// ユニットテストなどに使う.
func NewSQLiteDB(filePath string) *gorm.DB {
	os.Remove(filePath)

	db, err := gorm.Open(
		sqlite.Open(filePath),
		&gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Silent),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// sqliteは、デフォルトで外部キー制約を無視してしまうため
	if res := db.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(All...)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func NewTestDB() *gorm.DB {
	return NewSQLiteDB(":memory:")
}
