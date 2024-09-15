package main

import (
	"net/http"
)

func main() {
	r := initRouter()
	s := initServer(r)

	r.Get("/hello", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("world"))
	})

	app := &app{
		srv: s,
	}

	app.start()
}
