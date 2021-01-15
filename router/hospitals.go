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
	param, ok := r.URL.Query()["level"]
	if !ok  {
		fmt.Println("Missing GET pain level parameter")
	}
	var level int
	if len(param) < 1 {
		level = 1
	} else {
		level, _ = strconv.Atoi(param[0])
	}

	paths := []string{
		"./html/hospitals/hospitals.tmpl",
	}
	t := template.Must(template.New("hospitals.tmpl").ParseFiles(paths...))

	buf := new(bytes.Buffer)
	hospitals := middleware.Fetch_all_hospitals() //[]strcut.Hospital
	var wait_time []Wait_time = []Wait_time{}
	for _, v := range hospitals {
		wait_time = append(wait_time, Wait_time{v.Name, v.Waiting_list[level - 1].Patient_count * v.Waiting_list[level - 1].Average_process_time})
	}
	sort.Slice(wait_time, func(i, j int) bool {
		return wait_time[i].Wait < wait_time[j].Wait
	})

	err := t.Execute(buf, wait_time)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(buf.Bytes())
}

type Wait_time struct {
	Name string
	Wait int
}