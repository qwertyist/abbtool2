package main

import (
	"bytes"
	"time"
)

type ShortformResponse map[string][]string

type Abbreviation struct {
	ID      string `json:"id"`
	Abb     string `json:"abb"`
	Word    string `json:"word"`
	Creator string `json:"creator"`

	Comment string `json:"comment"`
	Remind  bool   `json:"remind"`
	ListID  string `json:"listId"`

	Updated time.Time `json:"updated"`
}

type protypeLists map[string]*bytes.Buffer

type textOnTopJSON struct {
	Autocorrect struct {
		Default struct {
			List map[string]string
		} `json:"<default>"`
	} `json:"autocorrect"`
	Shortform map[string]struct {
		Shortforms map[string]string
	}
}

type illumiTypeJSON struct {
	Lists []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"lists"`
	Abbreviations []struct {
		ListID       int    `json:"listId"`
		Abbreviation string `json:"abbreviation"`
		Word         string `json:"word"`
	} `json:"abbreviations"`
}
