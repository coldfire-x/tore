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

	// remove input
	re, _ = regexp.Compile("<input[\\S\\s]+?</input>")
	src = re.ReplaceAllString(src, "")

	// remove links
	re, _ = regexp.Compile("<link[\\S\\s]+?/>")
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

	output, err := iconv.ConvertString(src, encoding, "utf-8")
	// if error, return raw string
	if err != nil {
		log.Println(err)
		return src
	}

	return output
}
