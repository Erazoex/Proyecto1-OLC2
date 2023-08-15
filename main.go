package main

import "proyecto2/grammar"

func main() {
	grammar.Execute()
}

// antlr -Dlanguage=Go -o parser -visitor -no-listener Grammar.g4
