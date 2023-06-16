package config

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode   string
	HttpPort  string
	SecretKey string

	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Println("Configuration file reading error. Please check the file path: ", err)
		log.Println(err)
	}
	LoadServer(file)
	LoadMySQL(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("app").Key("APP_MODE").String()
	HttpPort = file.Section("app").Key("HTTP_PORT").String()
	SecretKey = file.Section("app").Key("SECRET_KEY").String()
}

func LoadMySQL(file *ini.File) {
	DbHost = file.Section("mysql").Key("DB_HOST").String()
	DbPort = file.Section("mysql").Key("DB_PORT").String()
	DbUser = file.Section("mysql").Key("DB_USER").String()
	DbPass = file.Section("mysql").Key("DB_PASS").String()
	DbName = file.Section("mysql").Key("DB_NAME").String()
}
