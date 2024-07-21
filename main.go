package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/somya4501/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is starting")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}
