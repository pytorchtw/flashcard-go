package models

type Deck struct {
	URL        string
	Title      string
	Flashcards []*Flashcard
	Content    string
}

type Flashcard struct {
	Title string
	Front string
	Back  string
}
