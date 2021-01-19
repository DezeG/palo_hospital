package router

import (
	"bytes"
	"text/template"
	"fmt"
	"../middleware"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		"./html/index/index.tmpl",
	}
	template_name := "index.tmpl"

	page := Construct_index(paths, template_name, r.Host)
	w.Write(page)
}

func Construct_index(paths []string, template_name string, host string) []byte {
	t := template.Must(template.New(template_name).ParseFiles(paths...))

	buf := new(bytes.Buffer)
	illnesses := middleware.Fetch_all_illnesses() //[]string
	templ := map[string]interface{} {
		"host": host,
		"illnesses": illnesses,
	}
	err := t.Execute(buf, templ)
	if err != nil {
		fmt.Println(err)
	}

	return buf.Bytes()
}