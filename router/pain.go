package router

import (
	"bytes"
	"text/template"
	"fmt"
	"net/http"
)

func pain(w http.ResponseWriter, r *http.Request) {
	param, ok := r.URL.Query()["illness"]
	if !ok  {
		fmt.Println("Missing GET illness parameter")
	}
	templ := map[string]interface{} {
		"illness": "",
		"host": r.Host,
	}

	if len(param) < 1 {
		templ["illness"] = "Unknown illness" 
	} else {
		templ["illness"] = param[0]
	}

	paths := []string{
		"./html/pain/pain.tmpl",
	}

	t := template.Must(template.New("pain.tmpl").ParseFiles(paths...))

	buf := new(bytes.Buffer)
	err := t.Execute(buf, templ)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(buf.Bytes())
}