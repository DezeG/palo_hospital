package middleware

import(
	"testing"
	"strings"
)

func TestDo_request(t *testing.T) {
	test := Do_request("https://wikipedia.org", "GET")
	if test == nil {
		t.Error("Did not retrieve a webpage")
	}
	if !(strings.Contains(string(test), "Wikipedia")) {
		t.Error("Did not fetch the page correctly")
	}
}

func TestFetch_all_hospitals(t *testing.T) {
	test := Fetch_all_hospitals()
	if test == nil {
		t.Error("Did not initiate hospital struct at all")
	}
	if len(test) < 1 {
		t.Error("Did not retrieve Azure API data")
	}
	for i := 0; i < len(test); i++ {
		if test[i].Id == 0 || test[i].Name == "" || test[i].Waiting_list == nil {
			t.Error("Missing fields in parced response")
		}
	}
}

func TestFetch_all_illnesses(t *testing.T) {
	test := Fetch_all_illnesses()
	if test == nil {
		t.Error("Did not initiate illness struct at all")
	}
	if len(test) < 1 {
		t.Error("Did not retrieve Azure API data")
	}
	for i := 0; i < len(test); i++ {
		if test[i] == "" {
			t.Error("Missing fields in parced response")
		}
	}
}