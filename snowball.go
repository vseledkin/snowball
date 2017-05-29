package snowball

import (
	"fmt"

	"github.com/vseledkin/snowball/english"
	"github.com/vseledkin/snowball/french"
	"github.com/vseledkin/snowball/russian"
	"github.com/vseledkin/snowball/spanish"
)

const (
	VERSION string = "v0.3.4"
)

// Stem a word in the specified language.
//
func Stem(word, language string, stemStopWords bool) (stemmed string, err error) {

	var f func(string, bool) string
	switch language {
	case "english":
		f = english.Stem
	case "spanish":
		f = spanish.Stem
	case "french":
		f = french.Stem
	case "russian":
		f = russian.Stem
	default:
		err = fmt.Errorf("Unknown language: %s", language)
		return
	}
	stemmed = f(word, stemStopWords)
	return

}

// IsStopWord test if a word is stop wort in the specified language.
//
func IsStopWord(word, language string) (ok bool, err error) {

	var f func(word string) bool
	switch language {
	case "english":
		f = english.IsStopWord
	case "spanish":
		f = spanish.IsStopWord
	case "french":
		f = french.IsStopWord
	case "russian":
		f = russian.IsStopWord
	default:
		err = fmt.Errorf("Unknown language: %s", language)
		return
	}
	return f(word), nil
}
