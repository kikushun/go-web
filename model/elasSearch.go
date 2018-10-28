package model

// SearchReq 検索リクエスト
type SearchReq struct {
	Query `json:"query"`
}

// Query query
type Query struct {
	Bool `json:"bool"`
}

// Bool bool
type Bool struct {
	Musts    []Must    `json:"must,omitempty"`
	Filters  []Filter  `json:"filter,omitempty"`
	Shoulds  []Should  `json:"should,omitempty"`
	MustNots []MustNot `json:"must_not,omitempty"`
}

// Must must
type Must struct {
	Term map[string]interface{} `json:"term,omitempty"`
}

// Filter filter
type Filter struct {
	Term map[string]interface{} `json:"term,omitempty"`
}

// Should should
type Should struct {
	Term map[string]interface{} `json:"term,omitempty"`
}

// MustNot mustnot
type MustNot struct {
	Term map[string]interface{} `json:"term,omitempty"`
}

// AddMust addMust
func (b *Bool) AddMust(searchType, field string, values interface{}) {

	fieldMap := map[string]interface{}{}
	fieldMap[field] = values

	switch searchType {
	case "term":
		b.Musts = append(b.Musts, Must{Term: fieldMap})
	}
}

// AddShould addShould
func (b *Bool) AddShould(searchType, field string, values interface{}) {

	fieldMap := map[string]interface{}{}
	fieldMap[field] = values

	switch searchType {
	case "term":
		b.Shoulds = append(b.Shoulds, Should{Term: fieldMap})
	}
}

// SearchResp ユーザ検索時のレスポンス
type SearchResp struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Hits     Hits `json:"hits"`
}

// Hits hits
type Hits struct {
	Total int   `json:"total"`
	Hits  []Doc `json:"hits"`
}

// Doc データ単体
type Doc struct {
	Index string `json:"_index"`
	ID    string `json:"_id"`
	User  User   `json:"_source"`
}
