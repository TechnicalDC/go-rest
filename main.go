/*
REST api in golang
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name		string
	City		string
	Dob		string
	Email 	string
	Password string
	Active	bool
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

	var user User;
	var users Users;

	user.Name = "Dilip Chauhan"
	user.City = "Mumbai"
	user.Active = true
	user.Email = "xyz@google.com"
	user.Password = "123"
	user.Dob = "01/01/9999"

	users.User = user
	data,_ := json.Marshal(users)

	io.WriteString(response, string(data))
}
