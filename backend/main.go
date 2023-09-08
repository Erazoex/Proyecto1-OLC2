package main

import (
	"proyecto2/api"
)

func main() {
	server := api.NewAPIServer(":4200")
	server.Run()
}

// antlr -Dlanguage=Go -o parser -visitor -no-listener Grammar.g4
