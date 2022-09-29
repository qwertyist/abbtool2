package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/yeka/zip"
)

func ImportTextOnTop(dat []byte) (ShortformResponse, error) {
	found := false
	var totlist textOnTopJSON
	var ac []string
	var imported = make(map[string][]string)

	err := json.Unmarshal(dat, &totlist)

	if err != nil {
		return nil, ErrUnmarshalTextOnTop
	}

	for abb, word := range totlist.Autocorrect.Default.List {
		fmt.Println(abb, "=", word)
		ac = append(ac, abb+"="+word)
	}

	if ac == nil {
		log.Println("Not a text-on-top export")
		return nil, ErrNotTextOnTopExport
	}

	for name, list := range totlist.Shortform {
		var abbs []string
		if name == "<default>" {
			name = "# Förkortningslista (Standard)"
		}
		for abb, word := range list.Shortforms {
			abb := abb + "=" + word
			abbs = append(abbs, abb)
		}
		if !found && len(abbs) > 0 {
			found = true
		}
		imported[name] = abbs
	}
	imported["# Autokorrigering (Standard)"] = ac
	if found {
		return imported, nil
	}
	return nil, ErrNoAbbs
}

func ImportProtype(buf []byte) (ShortformResponse, error) {
	found := false
	resp := make(map[string][]string)
	zipReader, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))

	if err != nil {
		log.Fatal(err)
	}

	for _, zipFile := range zipReader.File {
		if strings.HasSuffix(zipFile.Name, "wordlist.dat") {
			listName := zipFile.Name[:len(zipFile.Name)-12]
			list := toUTF8(listName)
			log.Println("[", list, "]")
			zipFile.SetPassword("SkrivTolk")
			r, err := zipFile.Open()
			if err != nil {
				return nil, fmt.Errorf("couldn't open wordlist.dat: %s", err.Error())
			}
			wl, err := ioutil.ReadAll(r)
			if err != nil {
				return nil, fmt.Errorf("couldn't read bytes of wordlist.dat: %s", err.Error())
			}
			abbs := ParseProtypeDAT(wl)
			if !found && len(abbs) > 0 {
				found = true
			}
			resp[list] = abbs
		}
	}

	if !found {
		return nil, ErrNoAbbs
	}

	return resp, nil
}

func ParseProtypeDAT(dat []byte) []string {
	rs := bytes.Runes(dat)
	var first bool
	first = true
	var abbs []string
	var rawAbb, rawWord []rune
	var length rune
	if len(rs) > 2 {
		length = rs[2]
	}
	log.Printf("First length: %d\n", length)
	log.Printf("File length: %d\n", len(rs[2:])+1)
	for i := range rs[2:] {
		if i+1 == len(rs[2:]) {
			abb := string(rawAbb) + "=" + string(rawWord)
			abbs = append(abbs, abb)
			rawWord = nil
			rawAbb = nil
			break

		}

		if length == 0 {
			length = rune(dat[i+3])
			if length == 255 {
				length += rune(dat[i+4]) + 3
			}
			//	log.Printf("Length is 0, restarting with %d\n", length)
			if first {
				first = false
			} else if first == false {
				first = true
				abb := string(rawAbb) + "=" + string(rawWord)
				abbs = append(abbs, abb)
				rawWord = nil
				rawAbb = nil
			}
			continue
		}
		if first == true {
			rawAbb = append(rawAbb, rune(dat[i+3]))
			length--
		} else {
			rawWord = append(rawWord, rune(dat[i+3]))
			length--
		}
	}
	return abbs
}

func ImportIllumiType(dat []byte) (ShortformResponse, error) {
	foundList := false
	foundAbb := false
	log.Println("Import illumitype lists")
	var illumiList illumiTypeJSON
	var listNames = make(map[int]string)
	var imported = make(map[string][]string)
	err := json.Unmarshal(dat, &illumiList)
	if err != nil {
		return nil, fmt.Errorf("ImportIllumiType|Couldn't unmarshal:\n%s", err.Error())
	}

	if len(illumiList.Lists) > 0 {
		foundList = true
	}

	for _, list := range illumiList.Lists {
		listNames[list.ID] = list.Name
	}
	for _, abb := range illumiList.Abbreviations {
		foundAbb = true
		log.Println(abb)
		a := abb.Abbreviation + "=" + abb.Word
		imported[listNames[abb.ListID]] = append(imported[listNames[abb.ListID]], a)
	}

	if !foundList {
		return nil, ErrNoLists
	}
	if !foundAbb {
		return nil, ErrNoAbbs
	}
	return imported, nil
}
