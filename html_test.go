package main

import (
	"testing"

	"github.com/Gabrieltrinidad0101/html-parser/lexer"
	"github.com/Gabrieltrinidad0101/html-parser/parser"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)

	lexer_, err := lexer.NewLexer("./index2.html")
	if err != nil {
		panic(err)
	}
	lexer_.Tokens()
	parser_ := parser.NewParser(lexer_.Targets)
	dom := parser_.Parser()

	testesOfH1111 := []string{
		".test_1 div #test_111",
		".test_1 #test_111",
		"#test_111",
	}

	for _, test := range testesOfH1111 {
		h1 := dom.QuerySelector(test)
		validateH1Test111(h1, t)
	}

	p := dom.QuerySelector("p")
	assert.Equal("Hello world", p.Children[0].TextContent)
	assert.Equal(map[string]string{}, p.Properties)
	assert.Equal("p", p.Type_)

	ps := dom.QuerySelectorAll("p")
	assert.Equal(2, len(*ps))

}

func validateH1Test111(test_111 *parser.Element, t *testing.T) {
	assert := assert.New(t)
	assert.Equal("h1", test_111.Children[0].TextContent)
	assert.Equal("test_111", test_111.Properties["id"])
	assert.Equal("test_111", test_111.Properties["class"])
	assert.Equal("red", test_111.Properties["color"])
	assert.Equal("h1", test_111.Type_)
}
