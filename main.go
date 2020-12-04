package main

import (
	"github.com/karldoenitz/Tigo/TigoWeb"
	"test_git_action/handlers"
)

var routers = []TigoWeb.Router{
	{"/hello", &handlers.HelloHandler{}, nil},
}

func main() {
	application := TigoWeb.Application{
		IPAddress:  "0.0.0.0",
		Port:       8888,
		UrlRouters: routers,
	}
	application.Run()
}
