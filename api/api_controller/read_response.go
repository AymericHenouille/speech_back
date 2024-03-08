package api_controller

type ReadResponse[T interface{}] struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Datas  []T `json:"datas"`
}

type RefreshResponse[T interface{}] struct {
	Time  string `json:"time"`
	Datas []T    `json:"datas"`
}
