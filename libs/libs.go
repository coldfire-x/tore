package libs

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
