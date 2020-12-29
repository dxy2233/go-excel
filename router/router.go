package router

import (
	"fmt"
	"net/http"
	"oppo-excel/utils"

	"github.com/gorilla/mux"
)

// Init 路由创建
func Init() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/upload", upload).Methods("POST", "OPTIONS")
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
