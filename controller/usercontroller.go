package controller

import (
	"github.com/joho/godotenv"
	conn "../connection"
	bitkub "../model"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"log"
)

func GetENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

////////////////////////////////////////////


type ReturnResponse struct {
     Status bool
     Result []User
     Updateat int64
}

func SayHello(w http.ResponseWriter, r *http.Request) {

	start := time.Now().UnixNano() / int64(time.Millisecond)

	data := FetchAllUsers()
	response := ReturnResponse{ Status: true, Result: data , Updateat: time.Now().Unix() }
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	end := time.Now().UnixNano() / int64(time.Millisecond)

	diff := end - start

	fmt.Println(diff)
}

////////////////////////////////////////////

type User struct {
	Id            	int
	AccountCode  	string
	TradingCredit 	float32
}

func FetchAllUsers() []User{
	GetENV()
	connection := bitkub.ConfigBitkubWeb()
	db := conn.GetConnectionMysql(connection)

	rows, err := db.Query("select id, account_code, trading_credit from users")

	if err != nil {
		return nil
	}

	var Id int
	var AccountCode string
	var TradingCredit float32
	var users []User

	for rows.Next() {
		err := rows.Scan(&Id, &AccountCode, &TradingCredit)
		if err != nil {
			return nil
		}

		users = append(users, User{ Id: Id, AccountCode: AccountCode, TradingCredit: TradingCredit } )
	}

	return users
}

////////////////////////////////////////////