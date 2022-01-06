package http

import "fmt"

//Handler - stores pointers to our comment service
type Handler struct {
}

//function is in capital letter because we want to call it from our main function
// Newhandler - returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

// Method
//SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
}
