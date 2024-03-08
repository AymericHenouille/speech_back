package api_controller

import (
	"net/http"
	"strconv"
)

const DEFAULT_LIMIT = 10
const DEFAULT_OFFSET = 0

func ExtractLimitOffset(r *http.Request) (limit, offset int) {
	values := r.URL.Query()
	if values.Has("limit") {
		limit = parseInt(values.Get("limit"))
	} else {
		limit = DEFAULT_LIMIT
	}

	if values.Has("offset") {
		offset = parseInt(values.Get("offset"))
	} else {
		offset = DEFAULT_OFFSET
	}
	return
}

func parseInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return result
}
