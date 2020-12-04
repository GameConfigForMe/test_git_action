package handlers

import "github.com/karldoenitz/Tigo/TigoWeb"

type HelloHandler struct {
	TigoWeb.BaseHandler
}

func (h *HelloHandler) Get() {
	h.ResponseAsJson(struct {
		Data string `json:"data"`
	}{"Hello World"})
}
