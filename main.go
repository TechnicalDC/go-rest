/*
REST api in golang
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	m "github.com/TechnicalDC/go-rest/models"
	r "github.com/TechnicalDC/go-rest/routers"
)

func main() {

	content, err := ioutil.ReadFile("./config.json")
	var payload m.Config

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	if json.Valid(content) {
		err = json.Unmarshal(content, &payload)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}

		r.Routers()
	} else {
		fmt.Println("Invalid JSON!")
	}
}
