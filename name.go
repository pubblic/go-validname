//+build !windows

package validname

import (
	"strings"
)

var Alt = AltMap{
	'/': '／', // Solidus -> Fullwidth Solidus
}

func (alt AltMap) Replace(s string) string {
	return strings.Map(func(r rune) rune {
		w := alt[r]
		if w != 0 {
			return w
		}
		return r
	}, s)
}

func RegularName(name string) (string, bool) {
	return Alt.Replace(name), true
}

func IsRegular(name string) bool {
	for _, r := range name {
		if Alt[r] != 0 {
			return false
		}
	}
	return true
}
