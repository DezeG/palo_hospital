package router

import (
	"bytes"
	"text/template"
	"fmt"
	"net/http"
)

func new_patient(w http.ResponseWriter, r *http.Request) {
	illness, ok := r.URL.Query()["illness"]
	if !ok  {
		fmt.Println("Missing GET illness parameter")
	}
	hospital, ok := r.URL.Query()["hospital"]
	if !ok  {
		fmt.Println("Missing GET hospital parameter")
	}
	pain_level, ok := r.URL.Query()["level"]
	if !ok  {
		fmt.Println("Missing GET pain level parameter")
	}

	templ := map[string]interface{} {
		"illness": illness[0],
		"hospital": hospital[0],
		"pain_level": pain_level[0],
	}

	paths := []string{
		"./html/new_patient/new_patient.tmpl",
	}

	t := template.Must(template.New("new_patient.tmpl").ParseFiles(paths...))

	buf := new(bytes.Buffer)
	err := t.Execute(buf, templ)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(buf.Bytes())
}