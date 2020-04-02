package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/pytorchtw/flashcard-go/models"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetURLJsonData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func LoadDeckFromURL(url string) (*models.Deck, error) {
	deck := models.Deck{}
	deck.URL = url
	jsonData, err := GetURLJsonData(deck.URL)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonData), &deck)
	if err != nil {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(deck.Content)
	if err != nil {
		return nil, err
	}

	deck.Content = string(decoded)
	return &deck, nil
}

func MakeFlashcards(content string) ([]*models.Flashcard, error) {
	lines := strings.Split(content, "\n")
	var flashcards []*models.Flashcard
	var frontBuf bytes.Buffer
	var backBuf bytes.Buffer
	for _, line := range lines {
		//line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "###") {
			// wrap up the previous flashcard
			if backBuf.Len() != 0 {
				flashcard := models.Flashcard{}
				flashcard.Front = frontBuf.String()
				flashcard.Back = backBuf.String()
				flashcards = append(flashcards, &flashcard)
				frontBuf.Reset()
				backBuf.Reset()
			}
			frontBuf.WriteString(line + "\n")
		} else {
			backBuf.WriteString(line + "\n")
		}
	}

	// last question
	flashcard := models.Flashcard{}
	flashcard.Front = frontBuf.String()
	flashcard.Back = backBuf.String()
	flashcards = append(flashcards, &flashcard)

	return flashcards, nil
}
