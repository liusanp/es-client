package models

type Bool struct {
	// must should
	Type string
	// 查询类型
	SearchType  string
	SearchField string
	SearchValue interface{}
}

type EsSearch struct {
	PageSize int
	CurrPage int
	Querys   []Bool
}
