package parser

import "html-parser/lexer"

type Element struct {
	lexer.Target
	Children []*Element
	Parent   *Element
}

func forEach(element *Element, cb func(*Element) bool) *Element {
	for _, child := range element.Children {
		stop := cb(child)
		if stop {
			return child
		}
		return forEach(child, cb)
	}
	return nil
}

func (e Element) GetElementById(id string) *Element {
	value := e.Properties["id"]

	if value == id {
		return &e
	}

	return forEach(&e, func(e *Element) bool {
		value := e.Properties["id"]
		return value == id
	})
}
