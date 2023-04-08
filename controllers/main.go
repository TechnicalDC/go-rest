package controllers

import (
	"fmt"
	"net/http"

	"github.com/TechnicalDC/go-rest/models"
)

func HomePage(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("<h1>Welcome</h1>"))
}

func CreateUser (reponse http.ResponseWriter, request *http.Request) {

}

func UpdateUser (reponse http.ResponseWriter, request *http.Request) {

}

func DeleteUser (reponse http.ResponseWriter, request *http.Request) {

}

func GetUser (reponse http.ResponseWriter, request *http.Request) {
	var user models.User
	var users models.Users

	user.Name = "Dilip Chauhan"
	user.Email = "xyz@gmail.com"
	user.City = "Mumbai"

	users.User[1] = user

	fmt.Println(user)
	fmt.Println(users)

	reponse.Header().Add("content-type","text/json")
	reponse.Write([]byte(user.Name))
}
