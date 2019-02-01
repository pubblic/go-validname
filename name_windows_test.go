package validate

import (
	"testing"
)

var tc = []struct {
	name        string
	regularName string
	regularized bool
}{
	{"", "", true},
	{"abc ", "abc", true},
	{"한글", "한글", true},
	{"a/b/c", "a／b／c", true},
	{".?.CON|", ".？.CON｜", true},
	{".PRN", ".PRN", true},
	{"PRN...  ", "PRN...  ", false},
}

func TestRegularName(t *testing.T) {
	for _, args := range tc {
		regularName, regularized := RegularName(args.name)
		if regularized != args.regularized {
			t.Errorf("%q is not correctly regularized", args.name)
			continue
		}
		if regularName != args.regularName {
			t.Errorf("expect %q, but actually got %q", args.regularName, regularName)
			continue
		}
	}
}
