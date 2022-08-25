package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", uploadFormHandler)
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

func uploadFormHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("upload form handler...")
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%s", h.Sum(nil))
		t, _ := template.ParseFiles("templates/upload.gtpl")
		t.Execute(w, token)
	} else {
		lists, err := checkUploadedFile(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Filen kunde inte laddas upp:" + err.Error()))
			return
		}
		t, err := template.ParseFiles("templates/response.gtpl")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Filen har ett felaktigt format, kontakta IT-support:" + err.Error()))
			return
		}

		t.Execute(w, lists)
	}
}

func checkUploadedFile(w http.ResponseWriter, r *http.Request) (ShortformResponse, error) {
	if err := r.ParseMultipartForm(20 * MB); err != nil {
		return nil, fmt.Errorf("couldn't parse multipart form: %s", err.Error())
	}

	r.Body = http.MaxBytesReader(w, r.Body, 20*MB)

	file, multiPartFileHeader, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("form file open failed: %s", err.Error())
	}

	if _, err := file.Seek(0, 0); err != nil {
		return nil, fmt.Errorf("couldn't set file position to start: %s", err.Error())
	}

	log.Printf("Name: %#v\n", multiPartFileHeader.Filename)
	fileType := multiPartFileHeader.Header.Get("Content-Type")
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}

	switch fileType {
	case "application/zip":
		resp, err := ImportProtype(buf.Bytes())
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return resp, nil
	}
	return nil, nil
}

func importHandler(w http.ResponseWriter, r *http.Request) {
}
