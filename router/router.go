package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Lent 测试
func Lent() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/index", someFunc).Methods("")
	r.HandleFunc("/index2", someFunc2)
	return r
}
func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("testestestest"))
}
func someFunc2(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("func2"))
}
