package tore

import (
	"github.com/pengfei-xue/tore/libs"
)

// get text from url
func GetText(url string) (string, string) {
	h := libs.Alg(url)
	h.RunAlg()
	return h.Title(), h.Text()
}

func SetAlg(name string) {
	libs.SetAlg(name)
}
