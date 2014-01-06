// http://www3.nd.edu/~tweninge/pubs/WH_TIR08.pdf

package libs

type TTR struct {
	Url, title, text, RawContent string
	ratio                        map[string]float32
}

// run algorithm
func (t *TTR) RunAlg() {
}

// remove scripts, stylesheets, input, and image
func (t *TTR) preprocess() {
}

// count text to tag ratio
func (t *TTR) countTextToTagRatio() {
}

// return text
func (t *TTR) Text() string {
	return "Text"
}

// return title
func (t *TTR) Title() string {
	return "Title"
}

func (t *TTR) Init(url string) {
	t.Url = url
}
