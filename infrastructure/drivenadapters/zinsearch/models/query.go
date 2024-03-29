package models

type Query struct {
	SearchType  string   `json:"search_type"`
	QueryString Term     `json:"query"`
	From        uint16   `json:"from"`
	Sort_fields []string `json:"sort_fields"`
	MaxResults  uint16   `json:"max_results"`
}
