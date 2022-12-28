/*
REST api in golang
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name		string	`json:"name"`
	City		string	`json:"city"`
	Dob		string	`json:"dateofbirth"`
	Email 	string	`json:"email"`
	Password string	`json:"-"`
	Active	bool  	`json:"active"`
}

type Users struct {
	User		User
}

type Config struct {
	Ip			string
	Port		string
	DbIp		string
	DbPort	string
	DbUser	string
	DbPass	string
}

func main() {
	content, err := ioutil.ReadFile("./config.json")
	var payload Config

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	if json.Valid(content) {
		err = json.Unmarshal(content, &payload)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
	} else {
		fmt.Println("Invalid JSON!")
	}

	handleRequest(payload)
}

func handleRequest(config Config) {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":" + config.Port, nil))
}

func homePage(response http.ResponseWriter, request *http.Request) {

	var users Users;

	user := User{
		City : "Mumbai",
		Name : "Dilip Chauhan",
		Active : true,
		Email : "xyz@google.com",
		Password : "123",
		Dob : "01/01/9999",
	}

	users.User = user

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(users)
}
