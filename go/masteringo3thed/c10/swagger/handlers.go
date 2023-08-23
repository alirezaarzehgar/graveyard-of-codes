// Package main for keeping handlers
//
// Documention for REST API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.7
//
//	Consumes:
//	- application/json
//
//	Produce:
//	- application/json
//
// swagger:meta
package main

import "net/http"

// User defines the structure for a Full User Record
//
// swagger:model
type User struct {
	// The ID for the user
	// in: body

	// required: false
	// min: 1
	ID int `json:"id"`

	// The username for user record
	// in: body

	// required: true
	Username string `json:"user"`

	// The password of user record
	// in: body

	// required: true
	Password string `json:"password"`
}

// swagger:route POST /login LoginUser
// Login user to application
//
// responses:
//
//	200: User
//	404: ErrorMessage
func LoginHandler(w http.ResponseWriter, r *http.Request) {

}
