package models

type EsSearch struct {
	PageSize    int                    `json:"pageSize"`
	CurrentPage int                    `json:"currentPage"`
	QueryJson   map[string]interface{} `json:"queryJson"`
	Index		string					`json:"index"`
}

// ResponseBody represents the structure of the outgoing response
type EsData struct {
	Total int			`json:"total"`
	Data  []interface{} `json:"data"`
}