package router

import (
	"bytes"
	"text/template"
	"fmt"
	"../middleware"
	"sort"
	"net/http"
	"strconv"
)

func hospitals(w http.ResponseWriter, r *http.Request) {
	l, ok := r.URL.Query()["level"]
	if !ok  {
		fmt.Println("Missing GET pain level parameter")
	}
	ill, ok := r.URL.Query()["illness"]
	if !ok  {
		fmt.Println("Missing GET illness parameter")
	}
	level, _ := strconv.Atoi(l[0])
	fmt.Println(ill, level)

	paths := []string{
		"./html/hospitals/hospitals.tmpl",
	}
	t := template.Must(template.New("hospitals.tmpl").ParseFiles(paths...))

	buf := new(bytes.Buffer)
	hospitals := middleware.Fetch_all_hospitals() //[]strcut.Hospital
	var wait_time []Wait_time = []Wait_time{}
	for _, v := range hospitals {
		wait_time = append(wait_time, Wait_time{v.Name, v.Waiting_list[level].Patient_count * v.Waiting_list[level].Average_process_time})
	}
	sort.Slice(wait_time, func(i, j int) bool {
		return wait_time[i].Wait < wait_time[j].Wait
	})

	templ := map[string]interface{} {
		"wait_time": wait_time,
		"illness": ill[0],
		"pain_level": level,
	}

	err := t.Execute(buf, templ)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(buf.Bytes())
}

type Wait_time struct {
	Name string
	Wait int
}