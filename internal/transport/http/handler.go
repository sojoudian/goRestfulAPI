package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sojoudian/goRestfulAPI/internal/comment"
)

//Handler - stores pointers to our comment service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

//function is in capital letter because we want to call it from our main function
// Newhandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// Method
//SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}
