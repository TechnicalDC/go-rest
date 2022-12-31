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
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	City		string	`json:"city"`
	Dob		string	`json:"dateofbirth"`
	Email 	string	`json:"email"`
	Password string	`json:"-"`
	Active	bool  	`json:"active"`
}

type Users struct {
	User		[]User
}

type Config struct {
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
	http.HandleFunc("/user", ManageUser)
	log.Fatal(http.ListenAndServe(":" + config.Port, nil))
}

func homePage(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("<h1>Welcome</h1>"))
}

func ManageUser(response http.ResponseWriter, request *http.Request) {

	var users Users;

	switch request.Method {
	case "GET":
		response.Header().Add("Content-Type", "application/json")
		json.NewEncoder(response).Encode(users)
	case "POST":
		var newuser User
		response.Header().Add("Content-Type", "application/json")

		json.NewDecoder(request.Body).Decode(&newuser)

		users.User = append(users.User, newuser)

		json.NewEncoder(response).Encode(users)
	}
}

