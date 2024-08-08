package main

import (
	"fmt"

	"github.com/Gabrieltrinidad0101/html-parser/lexer"
	"github.com/Gabrieltrinidad0101/html-parser/parser"
)

func main() {
	lexer_, err := lexer.NewLexer("./index.html")
	if err != nil {
		panic(err)
	}
	lexer_.Tokens()
	parser_ := parser.NewParser(lexer_.Targets)
	dom := parser_.Parser()
	dom.GetElementById("test")
	elemt := dom.QuerySelector(".1 #2 div > h1")
	fmt.Println(elemt)
}
