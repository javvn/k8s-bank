package utils

import "os"

//type APIConfig struct {
//	HeaderValue struct {
//		ClientID string
//	}
//}

//func NewAPIConfig() *APIConfig {
//	config := new(APIConfig)
//	config.HeaderValue.ClientID = os.Getenv("SBCNTR_CLIENT_ID_HEADER")
//
//	return config
//}

type ConfigDB struct {
	MySQL struct {
		DBMS     string
		Protocol string
		Username string
		Password string
		DBName   string
	}
}

func NewConfigDB() *ConfigDB {
	config := new(ConfigDB)

	config.MySQL.DBMS = "mysql"
	config.MySQL.Protocol = "tcp(" + os.Getenv("DB_HOST") + ":3306)"
	config.MySQL.Username = os.Getenv("DB_USERNAME")
	config.MySQL.Password = os.Getenv("DB_PASSWORD")
	config.MySQL.DBName = os.Getenv("DB_NAME")

	return config
}
