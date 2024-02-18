package models

type Response struct {
	Hits struct {
		Hits []Hit `json:"hits"`
	} `json:"hits"`
}

type Hit struct {
	ID     string `json:"_id"`
	Source Email  `json:"_source"`
}
