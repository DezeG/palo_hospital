package structs

type Illnesses struct {
	Embedded struct {
		Illnesses []struct {
			Illness struct {
				Name string `json:"name"`
				Id int `json:"id"`
			} `json:"illness"`
		} `json:"illnesses"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
	} `json:"_links"`
	Page struct {
		Size int  `json:"size"`
		Total_elements int `json:"totalElements"`
		Total_pages int `json:"totalPages"`
		Number int `json:"number"`
	} `json:"page"`
}

type Hospitals struct {
	Embedded struct {
		Hospitals []Hospital `json:"hospitals"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
	} `json:"_links"`
	Page struct {
		Size int  `json:"size"`
		Total_elements int `json:"totalElements"`
		Total_pages int `json:"totalPages"`
		Number int `json:"number"`
	} `json:"page"`
}

type Waiting_list struct {
	Patient_count int `json:"patientCount"`
	Level_of_pain int `json:"levelOfPain"`
	Average_process_time int `json:"averageProcessTime"`
}

type Hospital struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Waiting_list []Waiting_list `json:"waitingList"`
}







