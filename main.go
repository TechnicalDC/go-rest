/*
REST api in golang
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/TechnicalDC/go-rest/models"
	"github.com/TechnicalDC/go-rest/controllers"
)

func main() {
	fmt.Println("Hello World")

	content, err := ioutil.ReadFile("./config.json")
	var payload models.Config

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	if json.Valid(content) {
		err = json.Unmarshal(content, &payload)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
		fmt.Println(payload)
	} else {
		fmt.Println("Invalid JSON!")
	}
}
