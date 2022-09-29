package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/convert", convertHandler)
	r.HandleFunc("/favicon.ico", faviconHandler)
	r.PathPrefix("/styles").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("./static/styles"))))
	r.PathPrefix("/images").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images"))))
	http.Handle("/", accessControl(r))
	port := ":3434"
	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port:", port)
		errs <- http.ListenAndServe(port, nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("terminated: %s", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Id-Token, Cache-Control")

		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, req)
	})
}

const (
	MB = 1 << 20
)

func importHandler(w http.ResponseWriter, r *http.Request) {
}
