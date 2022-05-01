package model

type Date struct {
	FromDate string `json:"from_date"`
	ToDate string `json:"to_date"`
}

type Paginate struct {
	Page int `json:"page"`
	PerPage int `json:"per_page"`
}

type Pagination struct {
	Next    int `json:"next"`
	Perv    int `json:"perv"`
	Current int `json:"current"`
	Total   int `json:"total"`
}