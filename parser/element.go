package parser

import (
	"github.com/Gabrieltrinidad0101/html-parser/lexer"
)

type Element struct {
	lexer.Target
	Children []*Element
	Parent   *Element
	query    Query
}

func forEach(element *Element, cb func(*Element) bool) *Element {
	for _, child := range element.Children {
		stop := cb(child)
		if stop {
			return child
		}
		elemt := forEach(child, cb)
		if elemt != nil {
			return elemt
		}
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

func (e Element) querySelectorBase(textQuery string, getAll bool) []*Element {
	querys := e.query.Analyze(textQuery)
	currentElement := &e
	elements := []*Element{}

mainLoop:
	for i := 0; i < len(*querys); i++ {
		query := (*querys)[i]

		currentElement = forEach(currentElement, func(element *Element) bool {
			if query.TypeSearch != "element" {
				value := element.Properties[query.TypeSearch]
				if value == query.Search {
					return true
				}
			}
			return element.Type_ == query.Search
		})

		if query.SearchOnlySubChildren {
			i++
			query := (*querys)[i]
			for _, child := range currentElement.Children {
				if query.TypeSearch != "element" {
					value := child.Properties[query.TypeSearch]
					if value == query.Search {
						currentElement = child
						continue mainLoop
					}
				}
				if child.Type_ == query.Search {
					currentElement = child
					continue mainLoop
				}
			}
		}
		if currentElement != nil {
			elements = append(elements, currentElement)
		}
		if !getAll {
			return elements
		}
	}
	return elements
}

func (e Element) QuerySelector(textQuery string) *Element {
	elements := e.querySelectorBase(textQuery, false)
	return elements[0]
}

func (e Element) QuerySelectorAll(textQuery string) []*Element {
	elements := e.querySelectorBase(textQuery, false)
	return elements
}
