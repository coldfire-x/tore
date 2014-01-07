// http://www3.nd.edu/~tweninge/pubs/WH_TIR08.pdf

package libs

type TTR struct {
	Url, title, text, RawContent, cleaned string
	charset                               string
	ratio                                 map[string]float32
}

// run algorithm
func (t *TTR) RunAlg() {
	t.preprocess()
}

// remove scripts, stylesheets, input, and image
func (t *TTR) preprocess() {
	content, err := HTTP.Get(t.Url)
	if err != nil {
		panic(err)
	}

	// copy raw html
	t.RawContent = content

	// clean up html, remove style script and input tags
	t.cleaned = CleanUpHtml(content)

	// get charset
	t.charset = GetHtmlCharset(t.cleaned)
}

// count text to tag ratio
func (t *TTR) countTextToTagRatio() {
}

// return text
func (t *TTR) Text() string {
	return ConvertToUtf8(t.cleaned, t.charset)
}

// return title
func (t *TTR) Title() string {
	title := RetrieveTitleFromHtml(t.cleaned)
	return ConvertToUtf8(title, t.charset)
}

func (t *TTR) Init(url string) {
	t.Url = url
}
