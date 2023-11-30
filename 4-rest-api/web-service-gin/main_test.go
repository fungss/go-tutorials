package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAlbumsReturnsAllRecordsInJSON(t *testing.T) {
	client := http.Client{}
	resp, err := client.Get("http://localhost:8080/albums")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "application/json")
	var receviedData []album
	err = json.NewDecoder(resp.Body).Decode(&receviedData)
	if err != nil {
		t.Fatalf("Error decoding JSON: %v", err)
	}

	assert.NotEmpty(t, receviedData)
	assert.Equal(t, receviedData, albums)
}

func TestGetAlbumByID(t *testing.T) {
	client := http.Client{}
	resp, err := client.Get("http://localhost:8080/albums/2")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "application/json")
	var receviedData album
	err = json.NewDecoder(resp.Body).Decode(&receviedData)
	if err != nil {
		t.Fatalf("Error decoding JSON: %v", err)
	}
	var expectedData album
	for _, album := range albums {
		if album.ID == "2" {
			expectedData = album
		}
	}

	assert.NotEmpty(t, receviedData)
	assert.Equal(t, receviedData, expectedData)
}

func TestGetAlbumByID(t *testing.T) {
	client := http.Client{}
	resp, err := client.Get("http://localhost:8080/albums/2")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "application/json")
	var receviedData album
	err = json.NewDecoder(resp.Body).Decode(&receviedData)
	if err != nil {
		t.Fatalf("Error decoding JSON: %v", err)
	}
	var expectedData album
	for _, album := range albums {
		if album.ID == "2" {
			expectedData = album
		}
	}

	assert.NotEmpty(t, receviedData)
	assert.Equal(t, receviedData, expectedData)
}
