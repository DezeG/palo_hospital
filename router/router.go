package router

import(
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	index(w, r)
}

func Pain(w http.ResponseWriter, r *http.Request) {
	pain(w, r)
}

func Hospitals(w http.ResponseWriter, r *http.Request) {
	hospitals(w, r)
}

func New_patient(w http.ResponseWriter, r *http.Request) {
	new_patient(w, r)
}