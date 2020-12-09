package router

import (
	"fmt"
	"invoice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// Lent 路由创建
func Lent() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/index", upload).Methods("POST", "OPTIONS")
	r.Use(corsMiddleware)
	return r
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func upload(w http.ResponseWriter, r *http.Request) {
	file, header, _ := r.FormFile("file")
	defer file.Close()
	fileName := header.Filename
	f, _ := utils.ProcessedExcel(file)
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")
	if _, err := f.WriteTo(w); err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
