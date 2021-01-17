package middleware

import (
	"log"
	"io/ioutil"
	"net/http"
	"../env"
	"encoding/json"
	"../structs"
)

func Fetch_all_illnesses() []string {
	endpoint := env.Illnesses_endpoint
	method := env.Illnesses_method

	var illnesses_list []string
	for endpoint != "" {
		raw_body := Do_request(endpoint, method)

		var illnesses structs.Illnesses = structs.Illnesses{}
		Unmarsh_json(raw_body, &illnesses)

		endpoint = illnesses.Links.Next.Href
		for _, v := range illnesses.Embedded.Illnesses {
			illnesses_list = append(illnesses_list, v.Illness.Name)
		}
	}

	return illnesses_list

}

func Fetch_all_hospitals() []structs.Hospital {
	endpoint := env.Hospitals_endpoint
	method := env.Hospitals_method


	var hospitals_list []structs.Hospital
	for endpoint != "" {
		raw_body := Do_request(endpoint, method)

		var hospitals structs.Hospitals = structs.Hospitals{}
		Unmarsh_json(raw_body, &hospitals)

		endpoint = hospitals.Links.Next.Href
		hospitals_list = append(hospitals_list, hospitals.Embedded.Hospitals...)
	}

	return hospitals_list
}


func Do_request(endpoint string, method string) []byte {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	raw_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return raw_body
}

func Unmarsh_json(data []byte, mem interface{}) {
	err := json.Unmarshal(data, mem)
	if err != nil {
		log.Println(err)
	}
}










