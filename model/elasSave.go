package model

// SaveResp elasticsearchへinsert/updateしたときのレスポンス
type SaveResp struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
	Result  string `json:"result"`
}
