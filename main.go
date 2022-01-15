package main

import (
	MyRoute "AutoBiography/route"
	"net/http"
)

func main() {
	route := MyRoute.NewRoute()

	http.ListenAndServe(":3000", route)
}