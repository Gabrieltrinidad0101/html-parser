package parser

import "html-parser/lexer"

type Element struct {
	lexer.Target
	Children []*Element
	Parent   *Element
}

type parser struct {
	targets []*lexer.Target
}

func NewParser(targets []*lexer.Target) *parser {
	return &parser{
		targets: targets,
	}
}

func (p parser) Parser() *Element {
	dom := &Element{
		Target: lexer.Target{
			Type_:  "root",
			IsOpen: true,
		},
		Parent: nil,
	}
	currentState := dom
	for i, target := range p.targets {
		if currentState.Parent != nil && !p.targets[i].IsOpen {
			currentState = currentState.Parent
		}
		if !target.IsOpen {
			continue
		}
		newElement := &Element{
			Target: *target,
		}
		newElement.Parent = currentState
		currentState.Children = append(currentState.Children, newElement)
		currentState = newElement
	}
	return dom
}
