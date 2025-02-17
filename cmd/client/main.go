// Client for testing the server

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	endpoint := "http://localhost:8080/"
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, endpoint, http.NoBody)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status:", response.Status)
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Body:", string(body))
}
