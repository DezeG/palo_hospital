package main

import(
	"log"
	"fmt"
	"net/http"
	"./router"
)




func main() {
	http.HandleFunc("/", router.Index)
	http.HandleFunc("/pain", router.Pain)
	http.HandleFunc("/hospitals", router.Hospitals)
	http.HandleFunc("/new_patient", router.New_patient)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}
}