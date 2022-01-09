package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sojoudian/goRestfulAPI/internal/comment"
)

//GetComment - retrive comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		//fmt.Fprintf(w, "Unable to  parse UINT from ID")
		sendErrorResponse(w, "Unable to  parse UINT from ID", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		//fmt.Fprintf(w, "Error Retrieving comment by ID")
		sendErrorResponse(w, "Error Retrieving comment by ID", err)
		return
	}
	if err := sendOKResponse(w, comment); err != nil {
		// if err := json.NewEncoder(w).Encode(comment); err != nil {

		panic(err)
	}
	//fmt.Fprintf(w, "%+v", comment)
}

//GetAllComments - retrives all comments from the comment service
func (h *Handler) GetAllComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	comments, err := h.Service.GetAllComments()
	if err != nil {
		//fmt.Fprintf(w, "Fialed to retrive all comments")
		sendErrorResponse(w, "Fialed to retrive all comments", err)
		return
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
	//fmt.Fprintf(w, "%+v", comments)

}

// PostComment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		//fmt.Fprintf(w, "Failed to decode Json Body")
		sendErrorResponse(w, "Failed to decode Json Body", err)
		return
	}

	cmt, err := h.Service.PostComment(cmt)
	if err != nil {
		//fmt.Fprintf(w, "Failed to post new comment")
		sendErrorResponse(w, "Failed to post new comment", err)
		return
	}
	if err := sendOKResponse(w, cmt); err != nil {
		panic(err)
	}
	//fmt.Fprintf(w, "%+v", comments)
}

// UpdateComment - updates the comment by id
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		// fmt.Fprintf(w, "Failed to decode Json Body")
		sendErrorResponse(w, "Failed to decode Json Body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed from parse unit from ID", err)
		return
	}

	cmt, err = h.Service.UpdateComment(uint(commentID), cmt)
	if err != nil {
		sendErrorResponse(w, "Failed to update comment", err)
		return
	}
	if err := sendOKResponse(w, cmt); err != nil {
		panic(err)
	}
	//fmt.Fprintf(w, "%+v", comment)
}

//DeleteComment -delete comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		//fmt.Fprintf(w, "Failed to parse unit from ID")
		sendErrorResponse(w, "Failed to parse unit from ID", err)
		return
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		sendErrorResponse(w, "Failed to delete comment by comment ID", err)
		//fmt.Fprintf(w, "Failed to delete comment by comment ID")
		return
	}

	if err := sendOKResponse(w, Response{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
	// if err := json.NewEncoder(w).Encode(Response{Message: "Comment successfully deleted"}); err != nil {
	// 	panic(err)
	// }

	//fmt.Fprintf(w, "Successfully deleted comment")
}
