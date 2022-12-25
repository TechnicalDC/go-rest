// Rest api program
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
	User		[]User
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

	var reader []byte;

	body, err := request.Body.Read(reader)

	if err != nil {
		fmt.Println("Error")
	}

	// body,_ := io.ReadCloser(request.Body.Read(reader)).Read(reader)

	fmt.Println(reader, body)

	io.WriteString(response, "Hello, world!\n")
}
