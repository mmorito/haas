package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"

	_ "github.com/GoogleCloudPlatform/berglas/pkg/auto"
)

type env struct {
	// App
	Port        string `envconfig:"PORT" default:"8080"`
	Environment string `envconfig:"ENVIRONMENT" default:"local"` // local | ci | staging | production

	// GCP
	ProjectID      string `envconfig:"PROJECT_ID" default:"mnes-sandbox-morito"`
	Location       string `envconfig:"LOCATION" default:"asia-northeast1"`
	ServiceAccount string `envconfig:"SERVICE_ACCOUNT"`

	// PostgreSQL
	DBUserName            string `envconfig:"DB_USER_NAME"`
	DBPassword            string `envconfig:"DB_PASSWORD"`
	DBName                string `envconfig:"DB_NAME"`
	DBTcpHost             string `envconfig:"DB_TCP_HOST"`
	DBInstanceConnectName string `envconfig:"DB_INSTANCE_CONNECTION_NAME"`
}

var e *env

// SetEnv はenvをグローバル変数に設定する
func SetEnv() {
	var err error
	e, err = readFromEnv()
	if err != nil {
		log.Fatalf("failed to read env: %v", err)
	}
}

// Env はenvを取得する
func Env() *env {
	if e == nil {
		log.Fatalf("failed to get env")
	}
	return e
}

// readFromEnv は環境変数からenvに読み込む
func readFromEnv() (*env, error) {
	var e env
	if err := envconfig.Process("", &e); err != nil {
		return nil, err
	}
	return &e, nil
}
