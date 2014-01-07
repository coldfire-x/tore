// http://www3.nd.edu/~tweninge/pubs/WH_TIR08.pdf

package libs

import (
	"regexp"
	"strings"
)

type TTR struct {
	Url, title, text, RawContent, cleaned string
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

	// convert html tag to lowercase
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src := re.ReplaceAllStringFunc(content, strings.ToLower)

	// remove style
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	// remove script
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	// remove input
	re, _ = regexp.Compile("\\<input[\\S\\s]+?\\</input\\>")
	src = re.ReplaceAllString(src, "")

	// remove continuous line break
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	t.cleaned = src
}

// count text to tag ratio
func (t *TTR) countTextToTagRatio() {
}

// return text
func (t *TTR) Text() string {
	return t.cleaned
}

// return title
func (t *TTR) Title() string {
	if len(t.cleaned) < 1 {
		return ""
	}

	re, _ := regexp.Compile("\\<title([\\s\\S]+?)</title\\>")
	st := re.FindStringSubmatch(t.cleaned)
	if st == nil {
		return ""
	}

	title := st[1]
	if len(title) > 1 {
		return title[1:]
	}

	return ""
}

func (t *TTR) Init(url string) {
	t.Url = url
}
