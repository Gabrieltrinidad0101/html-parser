package parser

import (
	"maps"

	"github.com/Gabrieltrinidad0101/html-parser/lexer"
)

type parser struct {
	targets    []*lexer.Target
	cssAnalize *CssAnalize
}

func NewParser(targets []*lexer.Target) *parser {
	return &parser{
		targets:    targets,
		cssAnalize: NewCssAnalize(),
	}
}

func (p parser) Parser() *Element {
	csses := []string{}
	dom := NewElement(lexer.Target{
		Type_:  "root",
		IsOpen: true,
	})
	currentState := dom
	for i, target := range p.targets {

		if target.Type_ == "style" {
			csses = append(csses, target.TextContent)
			continue
		}

		if currentState.Parent != nil && !p.targets[i].IsOpen {
			currentState = currentState.Parent
		}

		if !target.IsOpen {
			continue
		}
		newElement := NewElement(*target)

		newElement.Parent = currentState
		currentState.Children = append(currentState.Children, newElement)
		if target.Type_ != "text" {
			currentState = newElement
		}
	}

	for _, css := range csses {
		queries := p.cssAnalize.Process(css)
		for query, properties := range queries {
			elements := dom.QuerySelectorAll(query)
			for _, element := range *elements {
				maps.Copy(element.Properties, properties)
			}
		}
	}

	return dom
}
