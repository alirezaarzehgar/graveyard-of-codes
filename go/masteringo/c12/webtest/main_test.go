package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/CheckStatusOK", nil)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckStatusOK)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned %v", rr.Code)
	}

	expect := "Fine!"
	if rr.Body.String() != expect {
		t.Errorf("Handler retured %v", rr.Body.String())
	}
}

func TestStatusNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/StatusNotFound", nil)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	StatusNotFound(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("handler returend %v", rr.Code)
	}

	if rr.Body.String() != "" {
		t.Errorf("handler returned %v", rr.Body.String())
	}
}

func TestMyHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/PATH", nil)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	MyHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returnred %v", rr.Code)
	}
	expect := "Serving: /PATH\n"
	expect += "Server: \n"
	if rr.Body.String() != expect {
		t.Errorf("handler returned %v", rr.Body.String())
	}
}
