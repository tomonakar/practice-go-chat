package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	GoogleClientSecret string
	GoogleClientId     string
	SecurityKey        string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("ファイルの読み込みに失敗しました: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		GoogleClientSecret: cfg.Section("google").Key("client_secret").String(),
		GoogleClientId:     cfg.Section("google").Key("client_id").String(),
		SecurityKey:        cfg.Section("securityKey").Key("key").String(),
	}
}
