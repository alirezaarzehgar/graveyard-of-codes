package main

// Generic error message returned as an HTTP Status Code
// swagger:response ErrorMessage
type ErrorMessage struct {
	// Description on situation
	// in: body
	Body int
}

// Generic noContent message returned as an HTTP Status Code
// swagger:response noContent
type NoContent struct {
	// Description on situation
	// in: body
	Body int
}

// Generic UserResponse message returned as an HTTP Status Code
// swagger:response User
type UserResponse struct {
	// Description on situation
	// in: body
	Body User
}
