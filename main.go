package main

import (
	"fmt"
	"html-parser/lexer"
	"html-parser/parser"
)

func main() {
	lexer_, err := lexer.NewLexer("./index.html")
	if err != nil {
		panic(err)
	}
	lexer_.Tokens()
	parser_ := parser.NewParser(lexer_.Targets)
	dom := parser_.Parser()
	fmt.Println(dom)
}
