package api

type ReadResponse[T interface{}] struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Data   T   `json:"data"`
}
