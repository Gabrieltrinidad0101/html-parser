package main

import (
	"fmt"

	"github.com/Gabrieltrinidad0101/html-parser/lexer"
	"github.com/Gabrieltrinidad0101/html-parser/parser"
)

func main() {
	lexer_, err := lexer.NewLexer("./index2.html")
	if err != nil {
		panic(err)
	}
	lexer_.Tokens()
	parser_ := parser.NewParser(lexer_.Targets)
	dom := parser_.Parser()
	h1s := dom.QuerySelector("div")
	for _, h1 := range h1s.Children {
		fmt.Println(h1.TextContent)
	}
}
