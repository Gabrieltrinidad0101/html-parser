package parser

type CssAnalize struct {
	idx         int
	css         string
	currentChar string
}

func NewCssAnalize() *CssAnalize {
	return &CssAnalize{
		idx: -1,
	}
}

func (ac *CssAnalize) advance() bool {
	if ac.idx >= len(ac.css)-1 {
		return false
	}
	ac.idx++
	ac.currentChar = string(ac.css[ac.idx])
	return true
}

func (ac *CssAnalize) skipSpace() bool {
	isSpace := false
	for ac.currentChar == " " || ac.currentChar == "\t" || ac.currentChar == "\r" || ac.currentChar == "\n" {
		canAdvance := ac.advance()
		if !canAdvance {
			return false
		}
		isSpace = true
	}
	return isSpace
}

func (ac *CssAnalize) Process(css string) map[string]map[string]string {
	ac.css = css
	var queries = map[string]map[string]string{}

main:
	for ac.advance() {
		query := ""
		for {
			addSpace := ac.skipSpace()
			if ac.currentChar == "{" {
				ac.advance()
				break
			}
			if addSpace && query != "" {
				query += " "
			}
			query += ac.currentChar
			canAdvance := ac.advance()
			if !canAdvance {
				break main
			}
		}
		ac.skipSpace()
		properties := map[string]string{}
		for ac.currentChar != "}" {
			var property string
			for ac.currentChar != ":" {
				property += ac.currentChar
				ac.advance()
				ac.skipSpace()
			}
			ac.advance()
			ac.skipSpace()
			var value string
			for ac.currentChar != ";" && ac.currentChar != "\n" {
				value += ac.currentChar
				ac.advance()
			}
			ac.advance()
			ac.skipSpace()
			properties[property] = value
		}
		queries[query] = properties
	}
	return queries
}
