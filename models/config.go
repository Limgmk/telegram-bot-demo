package models

import (
	"io/ioutil"
	"log"
	"tgb/utils"

	"gopkg.in/yaml.v2"
)

// Config :  配置文件结构
type Config struct {
	TelegramBotToken string `yaml:"bot-token"`
	Mode             string `yaml:"mode"`
	WebhookListener  WebhookListener `yaml:"webhook-listener"`
}

// Webhook-listener :  webhook 监听配置
type WebhookListener struct {
	TLS	bool `yaml:"tls"`
	SelfSigned bool `yaml:"self-signed"`
	CertPath string `yaml:"cert-path"`
	KeyPath	string `yaml:"key-path"`
	BindAddress string `yaml:"bind-address"`
	WebhookURL string `yaml:"webhook-url"`
}

var config *Config

func init() {
	// 解析配置文件
	config = new(Config)
	configBytes, _ := ioutil.ReadFile(utils.GetCurrentDirectory() + "/conf/config.yaml")

	if err := yaml.Unmarshal(configBytes, config); err != nil {
		log.Fatalln(err)
	}
}

func GetConfig() *Config {
	return config
}
