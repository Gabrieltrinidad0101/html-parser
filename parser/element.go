package parser

import (
	"time"

	"github.com/Gabrieltrinidad0101/html-parser/lexer"
)

type Element struct {
	lexer.Target
	Children []*Element
	Parent   *Element
	query    Query
}

func NewElement(target lexer.Target) *Element {
	return &Element{
		Target: target,
	}
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

func (e *Element) QuerySelector(textQuery string) *Element {
	queries := e.query.Analyze(textQuery)
	elements := make([]*Element, 0, len(queries))
	elements1 := e.querySelector(e, queries, 0, false, &elements)
	return (*elements1)[0]
}

func (e *Element) QuerySelectorAll(textQuery string) *[]*Element {
	queries := e.query.Analyze(textQuery)
	elements := make([]*Element, 0, len(queries))
	return e.querySelector(e, queries, 0, true, &elements)
}

func (e *Element) SetQueryFalses(queries []*QueryData) {
	for _, query := range queries {
		query.IsFound = false
	}
}

func (e *Element) querySelector(element *Element, queries []*QueryData, index int, getAll bool, elements *[]*Element) *[]*Element {
	for _, child := range element.Children {
		query := (queries)[index]

		if query.TypeSearch == "id" {
			query.IsFound = child.Properties["id"] == query.Search
		}

		if query.TypeSearch == "class" {
			query.IsFound = child.Properties["class"] == query.Search
		}

		if query.TypeSearch == "element" {
			query.IsFound = child.Type_ == query.Search
		}

		if query.IsFound {
			if index < len(queries)-1 {
				index++
			}
		}

		if (queries)[len(queries)-1].IsFound {
			*elements = append(*elements, child)
			e.SetQueryFalses(queries)
			if !getAll {
				return elements
			}
		}

		elemt := e.querySelector(child, queries, index, getAll, elements)
		query.IsFound = elemt != nil

		if !getAll && len(*elemt) == 1 {
			return elemt
		}
	}
	return elements
}
