// http://www3.nd.edu/~tweninge/pubs/WH_TIR08.pdf

package libs

import (
	"regexp"
	"strings"
    "math"
)

type TTR struct {
	Url, title, text, RawContent, cleaned string
	charset                               string
	ratio                                 map[int]float64
	sd                                    float64 // standard deviation
}

// run algorithm
func (t *TTR) RunAlg() {
	t.preprocess()
	t.countTextToTagRatio()
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
	// line no : ratio
	var tagratio = make(map[int]float64)

	lines := strings.Split(t.cleaned, "\n")
	for i, line := range lines {
		// get all chars in angle brackets
		re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
		matched := re.FindAllString(line, -1)

		// if no html tags found, ratio[i] = len(line)
		if matched == nil {
			tagratio[i] = float64(len(line))
			continue
		}

		// number of tags
		tags := len(matched)
		tagchars := 0
		for _, tt := range matched {
			tagchars += len(tt)
		}

		// number of non tags chars
		nontags := len(line) - tagchars

		// compute text/tag ratio
		tagratio[i] = float64(nontags) / float64(tags)
	}

	t.ratio = tagratio

	// adjust ratio
	radius := 2
	for i, _ := range t.ratio {
		// start from 3
		if i <= radius {
			continue
		}

		if i+radius > len(t.ratio) {
			break
		}

		// adjust ratio value, [i-radius, i+radius]
		var sum float64
		for j := i - radius; j < i+radius; j++ {
			sum += t.ratio[j]
		}

		t.ratio[i] = sum / (2.0*float64(radius) + 1.0)
	}

    var avg, sum float64
    for _, i := range t.ratio {
        sum += i
    }
    avg = sum / float64(len(t.ratio))

    sum = 0
    for _, i := range t.ratio {
        sum += math.Pow(i - avg, 2)
    }

    t.sd = math.Sqrt(sum/float64(len(t.ratio)-1))

	for i, m := range t.ratio {
        if m >= t.sd {
            t.text += lines[i]
        }
	}
}

// return text
func (t *TTR) Text() string {
	return ConvertToUtf8(t.text, t.charset)
}

// return title
func (t *TTR) Title() string {
	title := RetrieveTitleFromHtml(t.cleaned)
	return ConvertToUtf8(title, t.charset)
}

func (t *TTR) Init(url string) {
	t.Url = url
}
