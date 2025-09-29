package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Network requests")
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	// agent is required for this page
	req.Header.Set("User-Agent", "")

	response, err := client.Do(req)
	checkError(err)
	defer response.Body.Close()

	fmt.Printf("Response type: %T\n", response)

	bytes, err := io.ReadAll(response.Body)
	checkError(err)

	content := string(bytes)
	fmt.Print(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
