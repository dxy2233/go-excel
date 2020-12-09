package main

import (
	"net/http"

	"invoice/router"
)

func main() {
	r := router.Lent()
	http.ListenAndServe(":8999", r)
}
