package config

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Println("Configuration file reading error!")
		log.Println(err)
		return
	}
	LoadMySQL(file)
}

func LoadMySQL(file *ini.File) {
	DbHost = file.Section("mysql").Key("DB_HOST").String()
	DbPort = file.Section("mysql").Key("DB_PORT").String()
	DbUser = file.Section("mysql").Key("DB_USER").String()
	DbPass = file.Section("mysql").Key("DB_PASS").String()
	DbName = file.Section("mysql").Key("DB_NAME").String()
}
