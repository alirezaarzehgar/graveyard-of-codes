package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type City struct {
	Name string
	Area int
}

func shouldPost(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 - Method not allowed"))
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func justJSONContentType(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type"))
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func setServerTimeCookie(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		http.SetCookie(w, &http.Cookie{Name: "Server-Time", Value: strconv.FormatInt(time.Now().Unix(), 10)})
	}
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	var t City
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request Data"))
		return
	}

	log.Printf("%+v\n", t)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201 - Created"))
}

func main() {
	http.HandleFunc("/city", shouldPost(justJSONContentType(setServerTimeCookie(mainLogic))))
	http.ListenAndServe(":8000", nil)
}
