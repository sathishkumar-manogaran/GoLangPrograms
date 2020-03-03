package story_reading

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs string   `json:"paragraphs"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type Demo struct {
	Test string
	Name string
}
