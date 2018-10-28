package model

// Search 検索リクエスト
type Search struct {
	Query `json:"query"`
}

// Query query
type Query struct {
	Bool `json:"bool"`
}

// Bool bool
type Bool struct {
	Musts    []Must                   `json:"must,omitempty"`
	Filters  []map[string]interface{} `json:"Filter,omitempty"`
	Shoulds  []map[string]interface{} `json:"Should,omitempty"`
	MustNots []map[string]interface{} `json:"must_not,omitempty"`
}

// Must must
type Must struct {
	Term map[string]interface{} `json:"term,omitempty"`
}

// Filter filter
type Filter struct {
	Term map[string]interface{} `json:"filter,omitempty"`
}

// Should should
type Should struct {
	Term map[string]interface{} `json:"should,omitempty"`
}

// MustNot mustnot
type MustNot struct {
	Term map[string]interface{} `json:"must_not,omitempty"`
}

// AddMust addMust
func (s *Search) AddMust(searchType, field string, values interface{}) {

	fieldMap := map[string]interface{}{}
	fieldMap[field] = values

	switch searchType {
	case "term":
		s.Query.Bool.Musts = append(s.Query.Bool.Musts, Must{Term: fieldMap})
	}
}
