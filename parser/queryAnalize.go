package parser

import "strings"

type QueryData struct {
	TypeSearch            string
	SearchOnlySubChildren bool
	Search                string
}

type Query struct {
	query string
}

func NewQuery(query string) *Query {
	return &Query{
		query: query,
	}
}

func (q Query) Analyze(queryText string) *[]QueryData {
	querys := []QueryData{}
	querySplit := strings.Split(queryText, " ")
	for i := 0; i < len(querySplit); i++ {
		subQuery := querySplit[i]
		query := QueryData{}
		type_ := string(subQuery[0])
		if type_ == ">" {
			continue
		}
		if type_ == "#" {
			query.TypeSearch = "id"
		} else if type_ == "." {
			query.TypeSearch = "class"
		} else {
			query.TypeSearch = "element"
		}

		if query.TypeSearch != "element" {
			query.Search = subQuery[1:]
		} else {
			query.Search = subQuery
		}

		query.SearchOnlySubChildren = i < len(querySplit)-1 && string(querySplit[i+1]) == ">"
		querys = append(querys, query)
	}

	return &querys
}
