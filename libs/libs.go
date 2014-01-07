package libs

import (
	"log"
	"regexp"
	"strings"

	iconv "github.com/djimenez/iconv-go"
)

var (
	HTTP *httplib
)

// algrithom to extractor text from html
type OupengAlg interface {
	Init(url string)
	RunAlg()
	Text() string
	Title() string
}

func CleanUpHtml(src string) string {
	// convert html tag to lowercase
	re, _ := regexp.Compile("<[\\S\\s]+?>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// remove style
	re, _ = regexp.Compile("<style[\\S\\s]+?</style>")
	src = re.ReplaceAllString(src, "")

	// remove script
	re, _ = regexp.Compile("<script[\\S\\s]+?</script>")
	src = re.ReplaceAllString(src, "")

	// remove iframes
	re, _ = regexp.Compile("<iframe[\\S\\s]+?</iframe>")
	src = re.ReplaceAllString(src, "")

	// remove textare
	re, _ = regexp.Compile("<textarea[\\S\\s]+?</textarea>")
	src = re.ReplaceAllString(src, "")

	// remove input
	re, _ = regexp.Compile("<input[\\S\\s]+?</input>")
	src = re.ReplaceAllString(src, "")

	// remove comment
	re, _ = regexp.Compile("<!--[\\S\\s]+?-->")
	src = re.ReplaceAllString(src, "")

	// remove continuous line break
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return src
}

func GetHtmlCharset(src string) string {
	re, _ := regexp.Compile("<meta[\\s\\S]+?(charset=[-\\w]+)[\";\\s]*")
	st := re.FindStringSubmatch(src)
	if st == nil {
		return ""
	}

	charset := st[1]
	s := strings.SplitN(charset, "=", 2)[1]

	return s
}

func RetrieveTitleFromHtml(src string) string {
	if len(src) < 1 {
		return ""
	}

	re, _ := regexp.Compile("<title([\\s\\S]+?)</title>")
	st := re.FindStringSubmatch(src)
	if st == nil {
		return ""
	}

	title := st[1]
	if len(title) > 1 {
		return title[1:]
	}

	return ""
}

func ConvertToUtf8(src, encoding string) string {
	if len(src) < 1 {
		return src
	}

	if encoding == "utf-8" {
		return src
	}

	lines := strings.Split(src, "\n")
	var output string
	for _, line := range lines {
		o, err := iconv.ConvertString(line, encoding, "utf-8")

		// if error
		if err != nil {
			log.Println(err)
			output += line
			continue
		}

		output += o
	}

	return output
}
