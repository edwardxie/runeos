package service

import (
	// "fmt"
	// "html/template"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"

	// "runeos/ide"
)

func WebServer(port string) error {
	m := martini.Classic()
	m.Use(martini.Static("static/."))
	m.Use(render.Renderer(render.Options{
		Directory:  "views",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		// Funcs:      []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		Delims:     render.Delims{"{[{", "}]}"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))
	m.Get("/", Index)
	if err := http.ListenAndServe(port, m); err != nil {
		return err
	}
	return nil
}

func Index(r render.Render) {
	Data := make(map[string]string)
	Data["Website"] = "Runeos"
	Data["Author"] = "Edward Xie"
	Data["Email"] = "33066842@qq.com"
	Data["QQ"] = "33066842"
	r.HTML(200, "layout", Data)
}

func Webapp() {

}
