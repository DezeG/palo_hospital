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

	t := template.Must(template.New("index.tmpl").ParseFiles(paths...))

	buf := new(bytes.Buffer)
	illnesses := middleware.Fetch_all_illnesses() //[]string
	templ := map[string]interface{} {
		"host": r.Host,
		"illnesses": illnesses,
	}
	err := t.Execute(buf, templ)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(buf.Bytes())

}
