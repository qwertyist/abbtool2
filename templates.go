package main

import "fmt"

type tmplAssets struct {
	Title string
	Token string
	Lists ShortformResponse
}

var (
	ErrFormResend = fmt.Errorf("sidan förhindrade återsändning av formulär")
	ErrFormParse  = fmt.Errorf("filen är för stor eller fel format")
	ErrFileOpen   = fmt.Errorf("filen kunde inte öppnas")
	ErrFileSeek   = fmt.Errorf("filen har ett felaktigt format")
	ErrBufConvert = fmt.Errorf("internt fel, försök senare")

	ErrOpenProtypeWordlist = fmt.Errorf("kunde inte öppna wordlist.dat")
	ErrReadProtypeWordlist = fmt.Errorf("kunde inte läsa in wordlist.dat")
	ErrNoAbbs              = fmt.Errorf("inga förkortningar kunde importeras")
	ErrNoLists             = fmt.Errorf("inga listor kunde importeras")

	ErrUnmarshalIllumiType = fmt.Errorf("felaktigt .json-format")
	ErrUnmarshalTextOnTop  = fmt.Errorf("felaktigt .json-format")

	ErrNotTextOnTopExport = fmt.Errorf("filen ser inte ut som en TextOnTop-export")
)
