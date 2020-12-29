package main

import (
	"net/http"
	"oppo-excel/router"
)

func main() {
	r := router.Init()
	http.ListenAndServe(":8999", r)
}
