package config

import (
	"log"
	"squall/loghelper"
	"time"

	"github.com/caarlos0/env/v10"
)

var Config AppConfig

type AppConfig struct {
	// サーバを起動するホスト.
	Host string `env:"HOST" envDefault:"localhost"`
	// サーバを起動するポート.
	Port int `env:"PORT" envDefault:"8080"`
	// サーバを起動するアドレス.
	Address string `env:"ADDRESS,expand" envDefault:"$HOST:${PORT}"`

	// アプリケーションの動作モード
	AppEnv AppEnvType `env:"SQUALL_APP_ENV" envDefault:"development"`

	// "sqlite3" | "postgres"
	DBDriver string `env:"SQUALL_DB_DBDRIVER" envDefault:"sqlite3"`
	// DB接続情報
	DBConnectionStr string `env:"SQUALL_DB_DBCONNECTIONSTR" envDefault:"squall.sqlite"`
	// DB接続をリトライする回数.
	DBConnectionRetryAttempts uint `env:"SQUALL_DB_DBCONNECTION_RETRY_ATTEMPTS" envDefault:"9"`
	// DB接続をリトライする待ち時間.
	DBConnectionRetryDelay time.Duration `env:"SQUALL_DB_DBCONNECTION_RETRY_DELAY" envDefault:"100ms"`
	// テストなどで余計な標準出力がされないようにする.ログレベルを"error"にしても、異常系のエラーなどは表示されてしまうので、それらを一切表示しないために使う
	SilentMode bool `env:"SQUALL_SILENT_MODE" envDefault:"false"`
	// ログレベル。この設定レベル以下のログは表示・保存しない。
	// 多 <- "info" | "warn" | "error" -> 少
	LogLevel loghelper.LogLevel `env:"SQUALL_LOG_LEVEL" envDefault:"info"`
}

type AppEnvType string

var (
	AppEnvDevelopment AppEnvType = "development"
	AppEnvTesting     AppEnvType = "testing"
	AppEnvProduction  AppEnvType = "production"
)

func init() {
	if err := env.Parse(&Config); err != nil {
		log.Fatal(err)
	}
}

// GetDebugConfig はモックで設定を入れ替えるのに使う.
func GetDebugConfig(config *AppConfig) *AppConfig {
	return config
}
