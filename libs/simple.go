package libs

import (
	"regexp"
)

type Simple struct {
	Url, title, text, RawContent, cleaned string
	charset                               string
}

// run algorithm
func (s *Simple) RunAlg() {
	s.preprocess()
}

// remove scripts, stylesheets, input, and image
func (s *Simple) preprocess() {
	content, err := HTTP.Get(s.Url)
	if err != nil {
		panic(err)
	}

	// copy raw html
	s.RawContent = content

	// clean up html, remove style script and input tags
	s.cleaned = CleanUpHtml(content)

	// get charset
	s.charset = GetHtmlCharset(s.cleaned)
}

// return text
func (s *Simple) Text() string {
	re, _ := regexp.Compile("\\<img[\\S\\s]+?\\>")
	s.text = re.ReplaceAllString(s.cleaned, "")

	return ConvertToUtf8(s.text, s.charset)
}

// return title
func (s *Simple) Title() string {
	title := RetrieveTitleFromHtml(s.cleaned)
	return ConvertToUtf8(title, s.charset)
}

func (s *Simple) Init(url string) {
	s.Url = url
}
