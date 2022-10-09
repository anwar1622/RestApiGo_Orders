package main

import (
	"fmt"
	"log"
	"net/http"
	"orders_by/routers"
)

func main() {
	r := routers.Router()
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
