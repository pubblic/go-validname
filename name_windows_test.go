package validname

import (
	"testing"
)

func TestRegularName(t *testing.T) {
	tests := []struct {
		name    string
		regular string
		can     bool
	}{
		{"", "", true},
		{"abc ", "abc", true},
		{"한글", "한글", true},
		{"a/b/c", "a／b／c", true},
		{".?.CON|", ".？.CON｜", true},
		{".PRN", ".PRN", true},
		{"PRN...  ", "PRN...  ", false},
	}

	for _, test := range tests {
		regular, can := RegularName(test.name)
		if test.regular != regular {
			t.Errorf("RegularName(%q) is expected to return %q, not %q",
				test.name, test.regular, regular)
		}
		if test.can != can {
			t.Errorf("the name %q cannot be regularized",
				test.name)
		}
	}
}
