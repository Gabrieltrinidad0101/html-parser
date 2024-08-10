package lexer

import (
	"os"
	"strings"
)

type lexer struct {
	Text        string
	idx         int
	currentChar string
	Targets     []*Target
}

type Target struct {
	Type_       string
	IsOpen      bool
	Properties  map[string]string
	TextContent string
}

func NewLexer(path string) (*lexer, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &lexer{
		idx:  -1,
		Text: string(file),
	}, nil
}

func (l *lexer) Tokens() {
	l.advancer()
	for l.idx < len(l.Text)-1 {
		if l.currentChar == " " || l.currentChar == "\n" {
			l.advancer()
			continue
		}
		target := l.target()
		l.Targets = append(l.Targets, &target)
	}
}

func (l *lexer) advancer() {
	if l.idx >= len(l.Text)-1 {
		return
	}
	l.idx++
	l.currentChar = string(l.Text[l.idx])
}

func (l *lexer) target() Target {
	target := ""
	text := ""
	properties := map[string]string{}
	isOpen := true
	if l.currentChar == "<" {
		l.advancer()

		isOpen := l.currentChar != "/"

		if !isOpen {
			for l.currentChar != ">" {
				l.advancer()
			}
			l.advancer()
			return Target{}
		}

		for l.currentChar != " " && l.currentChar != ">" {
			target += l.currentChar
			l.advancer()
		}

		for l.currentChar == " " {
			l.advancer()
		}

	properties:
		for l.currentChar != ">" {
			propertyName := ""
			propertyValue := ""

			for l.currentChar != "=" {
				if l.currentChar == ">" {
					break properties
				}
				propertyName += l.currentChar
				l.advancer()
			}

			l.advancer()

			for l.currentChar != " " && l.currentChar != ">" {
				propertyValue += l.currentChar
				l.advancer()
			}
			if propertyName != "" {
				properties[strings.Trim(propertyName, " ")] = strings.Trim(propertyValue, "\"")
			}
		}
		l.advancer()
		for l.currentChar != "<" {
			text += l.currentChar
			l.advancer()
		}
	}

	return Target{
		Type_:       target,
		IsOpen:      isOpen,
		Properties:  properties,
		TextContent: text,
	}
}
