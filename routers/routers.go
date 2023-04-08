package routers

import (
	"log"
	"net/http"
)

func Routers() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/user", controllers.GetUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("<h1>Welcome</h1>"))
}
