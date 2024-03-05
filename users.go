package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const URI = "https://randomuser.me/api/?results=5000"

func makeMultipleApiCalls(numberOfCalls int) ([]Data, error) {
	var wg sync.WaitGroup

	// Create channels to receive data from goroutines
	usersChan := make(chan Data, numberOfCalls)
	errChan := make(chan error, numberOfCalls)

	for i := 0; i < numberOfCalls; i++ {
		wg.Add(1)

		go func() {

			defer wg.Done()

			// Make a GET request to the external API
			response, err := http.Get(URI)
			if err != nil {
				errChan <- err
				return
			}
			defer response.Body.Close()

			// Check if the request was successful (status code 200)
			if response.StatusCode != http.StatusOK {
				errChan <- fmt.Errorf("API request failed with status code: %d", response.StatusCode)
				return
			}

			// Decode the JSON response into a slice of User structs
			var users Data
			err = json.NewDecoder(response.Body).Decode(&users)
			if err != nil {
				errChan <- err
				return
			}

			usersChan <- users
		}()
	}

	// Close channels after the routines have finished
	go func() {
		wg.Wait()
		close(usersChan)
		close(errChan)
	}()

	// Get data from channels
	var usersData []Data
	for user := range usersChan {
		usersData = append(usersData, user)
	}

	// Check for errors from routines
	if err, ok := <-errChan; ok {
		return nil, err
	}

	return usersData, nil
}

// func getUsersFromExternalAPI() ([]Data, error) {

// 	// Make a GET request to the external API
// 	response, err := http.Get(URI)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer response.Body.Close()

// 	// Check if the request was successful (status code 200)
// 	if response.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("API request failed with status code: %d", response.StatusCode)
// 	}

// 	// Decode the JSON response into a slice of User structs
// 	var users []Data
// 	err = json.NewDecoder(response.Body).Decode(&users)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }
