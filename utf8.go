package main

import (
	"log"

	"golang.org/x/text/encoding/charmap"
)

func toUTF8(latin1 string) string {
	log.Println("decode latin1 string [", string(latin1), "] to utf-8 ( len: ", len(latin1), ")")
	d := charmap.ISO8859_1.NewEncoder()
	out, err := d.String(latin1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out)
	return out
}
