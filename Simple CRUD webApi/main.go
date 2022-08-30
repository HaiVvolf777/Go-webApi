package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/haider-star/mongodb/router"
)

func main() {
	fmt.Println("MongoDb Api connection")

	r := router.Router()
	fmt.Println("Server is getting Started On ......")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listing At Port 4000....")
}
