package connection

import (
    "log"
	"fmt"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type ConfigDatabase struct {
	Driver 		string
	Username 	string
	Password 	string
	Host 		string
	Port 		string
	Name 		string
}

func SetConfig(config []string) ConfigDatabase{
	return ConfigDatabase{ Driver: config[0], Username: config[1], Password: config[2], Host: config[3], Port: config[4], Name: config[5]}
}


// https://github.com/go-sql-driver/mysql

func GetConnectionMysql(dbConnection []string) (*sql.DB){

	configDB := SetConfig(dbConnection)
	fmt.Println(configDB.Driver, configDB.Username, configDB.Password, configDB.Host, configDB.Password, configDB.Name)

	fmt.Println("Connection...")
	db, err := sql.Open( configDB.Driver, configDB.Username + ":" + configDB.Password + "@tcp(" + configDB.Host + ":" + configDB.Port + ")/" + configDB.Name + "?timeout=5s" )
	
	if err != nil {
		fmt.Println("Connection Fail 01")
		return nil
	}
	if err = db.Ping(); err != nil {
		fmt.Println("Connection Fail 02")
		return nil
	}

	fmt.Println("Connection Success")
	return db
}

func GetENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}