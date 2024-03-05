package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := makeMultipleApiCalls(3)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching users: %s", err), http.StatusInternalServerError)
		return
	}

	unduplicatedUsers := removeDuplicatedUsers(users)

	jsonData, err := json.Marshal(unduplicatedUsers)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)

}

func removeDuplicatedUsers(data []Data) map[string]Results {
	u := make(map[string]Results)
	i := 0
	for _, v := range data {
		for _, y := range v.Results {
			if _, value := u[y.Login.UUID]; !value {
				i++
				u[y.Login.UUID] = y
			}
		}
	}

	return u
}
