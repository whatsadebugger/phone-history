package main

import (
	"io"
	"net/http"
)

// BasicPost will make a POST request to given url and with given body
func BasicPost(url string, contentType string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", contentType)

	resp, err := client.Do(request)
	return resp, err
}

// BasicGet will make a Get request to given url
func BasicGet(url string) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(request)
	return resp, err
}

// BasicDelete will make a DELETE request to given url and with given body
func BasicDelete(url string, contentType string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("DELETE", url, body)
	request.Header.Set("Content-Type", contentType)

	resp, err := client.Do(request)
	return resp, err
}

// BasicPut will make a put request to given url and with given body
func BasicPut(url string, contentType string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("PUT", url, body)
	request.Header.Set("Content-Type", contentType)

	resp, err := client.Do(request)
	return resp, err
}

// PostTextCSV post a csv file to a url
func PostTextCSV(url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", "text/csv")

	resp, err := client.Do(request)
	return resp, err
}
