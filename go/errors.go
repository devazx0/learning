package main

import (
	"fmt"
)

type NetworkResponse struct {
	errMessage,
	successMessage string
	statusCode int
}

func (e NetworkResponse) Error() error {
	return fmt.Errorf("%s with status code %d", e.errMessage, e.statusCode)
}

func (e NetworkResponse) Success() string {
	return fmt.Sprintf("%s with status code %d", e.successMessage, e.statusCode)
}

func pingNetwork(url string, success bool) (string, error) {
	if !success {
		response := NetworkResponse{
			errMessage:     "Not found",
			successMessage: "",
			statusCode:     404,
		}
		return "", response.Error()

	}

	response := NetworkResponse{
		successMessage: "URL is healthy",
		errMessage:     "",
		statusCode:     200,
	}

	return response.Success(), nil
}

func main() {

	ping, error := pingNetwork("test.com", true)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(ping)

}
