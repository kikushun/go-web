package model

// ElasResp elasticsearchへinsert/update/deleteしたときのレスポンス
type ElasResp struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
	Result  string `json:"result"`
}
