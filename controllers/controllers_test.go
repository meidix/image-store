package controllers

import (
	"mediastore/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testServerAddress = "localhost:4000"

func TestIndexEndpoint(t *testing.T) {
	router := router.Router()
	router.GET("/", HelloWorld)
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Errorf(`The returned status code is %v instead of 200`, response.Code)
	}
}