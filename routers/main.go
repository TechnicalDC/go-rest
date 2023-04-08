package routers

import (
	"log"
	"net/http"
	c "github.com/TechnicalDC/go-rest/controllers"
)

func Routers() {
	http.HandleFunc("/", c.HomePage)
	http.HandleFunc("/user", c.GetUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

