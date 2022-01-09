package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sojoudian/goRestfulAPI/internal/comment"
)

//Handler - stores pointers to our comment service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - an object to strore response from our API
type Response struct {
	Message string
	Error   string
}

//function is in capital letter because we want to call it from our main function
// Newhandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

//LoggingMiddleware - add middleware around endpoints
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"Method": r.Method,
			"Path":   r.URL.Path,
		}).Info("Handled request")
		// log.Info("Endpoint hit")
		next.ServeHTTP(w, r)
	})
}

// BasicAuth - a handly middleware function that will provide basic authentication around specific endpoint -
func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// log.WithFields(log.Fields{
		// 	"Method": r.Method,
		// 	"Path":   r.URL.Path,
		// }).Info("Basic aut endpoint hit")
		log.Info("Basic aut endpoint hit")
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "admin" && ok {
			original(w, r)
		} else {
			// w.Header().Set("Content-Type", "application/json; charset=utf8")
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
	}
}

// validateToken -
func validateToken(accessToken string) bool {
	var mySingingKey = []byte("missionimpossible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There has been an error")
		}
		return mySingingKey, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

//JWTAuth - decorator function for jwt validation for endpoints
func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("JWT authentication hit")
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			// w.Header().Set("Content-Type", "application/json; charset=utf-8")
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}

		// Bearer JWT-token
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			// w.Header().Set("Content-Type", "application/; charset=utf-8")
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}

		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			// w.Header().Set("Content-Type", "application/json; charset=utf8")
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
	}
}

// Method
//SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	log.Info("Setting up Routes")
	// fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment", h.GetAllComment).Methods("GET")
	h.Router.HandleFunc("/api/comment", JWTAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "I am alive!")
		// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// w.WriteHeader(http.StatusOK)
		if err := sendOKResponse(w, (Response{Message: "I am alive!"})); err != nil {
			panic(err)
		}

	})
}

//Define a way to set headers only once  for all endpoints
func sendOKResponse(w http.ResponseWriter, resp interface{}) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		log.Error(err)
	}
}
