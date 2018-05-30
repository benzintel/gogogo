package model

import (
	"os"
)

func ConfigMongoTest() []string{ 

	connection := []string{
		os.Getenv("CONNECTION_WEB_DRIVER"), 
		os.Getenv("CONNECTION_WEB_USERNAME"),
		os.Getenv("CONNECTION_WEB_PASSOWRD"),
		os.Getenv("CONNECTION_WEB_HOST"),
		os.Getenv("CONNECTION_WEB_PORT"),
		os.Getenv("CONNECTION_WEB_DBNAME"),
	}

	return connection
}