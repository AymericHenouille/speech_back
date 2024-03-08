package api_controller

import "net/http"

type ApiController struct {
	Get    func(w http.ResponseWriter, r *http.Request)
	Post   func(w http.ResponseWriter, r *http.Request)
	Put    func(w http.ResponseWriter, r *http.Request)
	Delete func(w http.ResponseWriter, r *http.Request)
}

func (c *ApiController) HandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			callHandler(c.Get, w, r)
		case http.MethodPost:
			callHandler(c.Post, w, r)
		case http.MethodPut:
			callHandler(c.Put, w, r)
		case http.MethodDelete:
			callHandler(c.Delete, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func callHandler(handler func(w http.ResponseWriter, r *http.Request), w http.ResponseWriter, r *http.Request) {
	if handler != nil {
		handler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
