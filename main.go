package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = 8080

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", UsersHandler).Methods("GET")

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), r)
	if err != nil {
		fmt.Println("Error starting http server: ", err)
	}
}
