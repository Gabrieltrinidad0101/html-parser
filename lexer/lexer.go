package lexer

import "os"

type lexer struct {
	Text        string
	idx         int
	currentChar string
	Targets     []*Target
}

type Property struct {
	Name  string
	Value string
}

type Target struct {
	Type_      string
	IsOpen     bool
	Properties []Property
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
	properties := []Property{}
	isOpen := true
	if l.currentChar == "<" {
		l.advancer()

		if l.currentChar == "/" {
			isOpen = false
		}

		for l.currentChar != " " && l.currentChar != ">" {
			target += l.currentChar
			l.advancer()
		}

		for !isOpen && l.currentChar != ">" {
			l.advancer()
		}

		for isOpen && l.currentChar != ">" {
			propertyName := ""
			propertyValue := ""
			for l.currentChar != "=" {
				propertyName += l.currentChar
				l.advancer()
			}

			for l.currentChar != " " {
				propertyValue += l.currentChar
				l.advancer()
			}
			if propertyName != "" {
				properties = append(properties, Property{
					Name:  propertyName,
					Value: propertyValue,
				})
			}
			l.advancer()
		}
		l.advancer()
	}
	return Target{
		Type_:      target,
		IsOpen:     isOpen,
		Properties: properties,
	}
}
