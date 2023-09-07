package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/grammar", makeHTTPHandleFunc(s.handleGrammar))

	log.Println("API server running on port:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func createAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) handleGrammar(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetGrammar(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetGrammar(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetGrammar(w, r)
	}
	return nil
}

func helloWorld() {
	fmt.Println("hello, world!")
}
