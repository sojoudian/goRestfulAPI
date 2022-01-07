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
	
	h.Router.HandleFunc("/api/comment", h.GetAllComment).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}

//GetComment - retrive comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	var := mux.Vars(r)
	id := vars["id"]

	i, er := strconv.ParseInt(id, 10, 64)
	if er != nil {
		fmt.Fprintf(w, "Unable to  parse UINT from ID")
	}

	comment, err := h.Service.GetComment(uint(id))
	if err != nil {
		fmt.Fprintf(w, "Error Retrieving comment by ID")
	}

	fmt.Fprintf(w, "%+v", comment)
}

//GetAllComments - retrives all comments from the comment service
func (h *Handler) GetAllComment(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "Fialed to retrive all comments")
	}
	fmt.Fprintf(w, "%+v", comments)

}

// PostComment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.PostComment(comment.Comment{
		Slug: "/"
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}
	fmt.Fprintf(w, "%+v", comments)
}

// UpdateComment - updates the comment by id
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request){
	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new"
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to update comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

//DeleteComment -delete comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse unit from ID")
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Fprintf(w, "Failed to delete comment by comment ID")
	}

	fmt.Fprintf(w, "Successfully deleted comment")
}