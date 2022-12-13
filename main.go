package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}

func serve(w http.ResponseWriter, r *http.Request) {
	var filename string

	if strings.Contains(r.URL.Path, ".js") {
		filename = fmt.Sprintf("./static%s", r.URL.Path)
		w.Header().Set("Content-Type", "text/javascript")
	} else if strings.Contains(r.URL.Path, ".css") {
		filename = fmt.Sprintf("./static%s", r.URL.Path)
		w.Header().Set("Content-Type", "text/css")
	} else if strings.Contains(r.URL.Path, ".ico") {
		filename = fmt.Sprintf("./static%s", r.URL.Path)
		w.Header().Set("Content-Type", "image/x-icon")
	} else {
		filename = fmt.Sprintf("./static/index.html")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	http.ServeFile(w, r, filename)
}

func main() {
	port := getPort()
	router := mux.NewRouter()

	router.PathPrefix("/").HandlerFunc(serve)

	fmt.Println("Server listening on port " + port)
	http.ListenAndServe(port, router)
}
