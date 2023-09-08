package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"proyecto2/grammar"
	"proyecto2/parser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain")
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

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/gramatica", makeHTTPHandleFunc(s.handleGrammar))

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
	input := r.URL.Query().Get("gramatica")
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewGrammarLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGrammarParser(tokenStream)
	p.BuildParseTrees = true
	tree := p.Init()
	eval := grammar.Visitor{
		Environment: grammar.NewEnvironment(nil),
	}
	eval.Visit(tree)
	cst := grammar.CreateCST(tree)
	err := os.WriteFile("cstree.dot", []byte(cst), 0644)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("dot", "-Tpdf", "cstree.dot", "-o", "cstree.pdf")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error al convertir el archivo .dot a .pdf", err.Error())
		log.Fatal(err)
	}
	grammar.GenerateHTML(&eval)
	grammar.GenerateHTMLERROR(&eval)
	fmt.Fprint(w, eval.GetConsole())
	return nil
}
