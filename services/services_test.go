package services

import (
	"github.com/pytorchtw/flashcard-go/models"
	"github.com/pytorchtw/flashcard-go/utils"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		//fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		log.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

type TestManager struct {
	Deck *models.Deck
}

var test TestManager

func Test_setup(t *testing.T) {
	test = TestManager{}
	test.Deck = &models.Deck{}
	data, err := ioutil.ReadFile(utils.Basepath + "/testdata/test_flashcards.md")
	ok(t, err)
	test.Deck.Content = string(data)
}

func Test_LoadDeckFromURL(t *testing.T) {
	url := "https://api.github.com/repos/pytorchtw/flashcard-go/contents/README.md"
	deck, err := LoadDeckFromURL(url)
	ok(t, err)
	equals(t, "# flashcard-go", deck.Content)
}

func Test_LoadFlashcards(t *testing.T) {
	flashcards, err := MakeFlashcards(test.Deck.Content)
	ok(t, err)
	equals(t, 3, len(flashcards))
}
