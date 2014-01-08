package tore

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
)

var tpl string = `
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html;charset=utf-8">
  <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css">
  <script src="//netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
  <style type="text/css">
  .container {padding-top: 30px;}
  </style>
</head>
<body>
<div class="container">
  <div class="jumbotron">
    <h2>输入URL, 提取正文</h2>
      <form method="post" action="/">
      <div class="input-group">
        <input class="form-control" name="q" autofocus="autofocus" value="" placeholder="url" type="text">
        <span class="input-group-btn">
        <button class="btn btn-default" type="submit">Go!</button>
        </span>
      </div>
      </form>
  </div>
  <div class="content">
  <h3>{{ .Title }}</h3>
  {{ .Text }}
  </div>
</div>
</body>
</html>`

type Page struct {
	Title string
	Text  template.HTML
}

func getFormUrl(u string) (string, error) {
	r, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	if r.Scheme == "" {
		r.Scheme = "http"
	}

	return r.String(), nil
}

// handler http request
func NewHttpHandler() http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rev := recover(); rev != nil {
				w.WriteHeader(500)
				log.Println(rev)
				debug.PrintStack()
			}
		}()

		if r.URL.Path != "/" {
			w.WriteHeader(403)
			return
		}

		var err error
		tmpl, _ := template.New("index").Parse(tpl)

		if r.Method == "GET" {
			err = tmpl.Execute(w, nil)
			if err != nil {
				panic(err)
			}
		} else if r.Method == "POST" {
			u := r.FormValue("q")
			log.Println(u)

			u, err = getFormUrl(u)
			if err != nil {
				err = tmpl.Execute(w, "Invalid URL")
				if err != nil {
					panic(err)
				}
				return
			}

			title, text := GetText(u)
			data := &Page{title, template.HTML(text)}
			err = tmpl.Execute(w, data)
			if err != nil {
				panic(err)
			}
		} else {
			w.WriteHeader(403)
		}
	}
	return http.HandlerFunc(h)
}
