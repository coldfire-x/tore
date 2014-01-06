package tore

import (
	"github.com/pengfei-xue/tore/libs"
)

var Alg libs.OupengAlg

// get text from url
func GetText(url string) (string, string) {
	Alg.Init(url)
	Alg.RunAlg()
	return Alg.Title(), Alg.Text()
}

func SetAlg(name string) {
	switch name {
	default:
		Alg = &libs.TTR{}
	}
}

func init() {
	Alg = &libs.TTR{}
}
