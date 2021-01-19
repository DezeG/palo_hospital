package router

import(
	"testing"
	"strings"
)

func TestConstruct_index(t *testing.T) {
	style := "stylesheet"
	title := "Hospital Assistant"
	test := Construct_index([]string{"../html/index/index.tmpl"}, "index.tmpl", "")
	if test == nil {
		t.Error("Did not create page at all")
	}
	if len(test) == 0 {
		t.Error("The page is empty")
	}
	if !(strings.Contains(string(test), title)) {
		t.Error("Missing title in the header")
	}
	if !(strings.Contains(string(test), style)) {
		t.Error("Missing style in the header")
	}
}