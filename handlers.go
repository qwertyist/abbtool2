package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	assets := tmplAssets{
		Title: "Start",
	}
	log.Println("Show start page")

	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	assets.Token = fmt.Sprintf("%s", h.Sum(nil))

	t, err := template.ParseFiles("templates/base.gtpl", "templates/main.gtpl")
	if err != nil {
		log.Println("couldn't create template:", err)
	}
	t.Execute(w, assets)
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	assets := tmplAssets{
		Title: "Konvertera förkortningslistor",
	}
	if r.Method == "GET" {
		log.Println("prevent form resend")
		w.Write([]byte("Du behöver ladda upp filen på nytt"))
		return
	} else {
		log.Println("upload form handler...")
		var err error
		assets.Lists, err = checkUploadedFile(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Filen kunde inte laddas upp:" + err.Error()))
			return
		}
		t, err := template.ParseFiles("templates/base.gtpl", "templates/response.gtpl")
		if err != nil {
			return
		}

		t.Execute(w, assets)
	}
}

func checkUploadedFile(w http.ResponseWriter, r *http.Request) (ShortformResponse, error) {
	if err := r.ParseMultipartForm(20 * MB); err != nil {
		return nil, fmt.Errorf("couldn't parse multipart form: %s", err.Error())
	}

	r.Body = http.MaxBytesReader(w, r.Body, 20*MB)

	file, multiPartFileHeader, err := r.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("form file open failed: %s", err.Error())
	}
	defer file.Close()

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
	case "application/json":
		resp, err := ImportTextOnTop(buf.Bytes())
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		if len(resp) > 0 {
			return resp, nil
		} else {
			resp, err = ImportIllumiType(buf.Bytes())
			if err != nil {
				return nil, err
			}
			return resp, nil
		}
	}
	return nil, nil
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon.ico")
}
