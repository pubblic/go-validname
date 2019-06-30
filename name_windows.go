//+build windows
package validname

import (
	"strings"
	"unicode"
)

// https://stackoverflow.com/a/31976060
// https://msdn.microsoft.com/en-us/library/windows/desktop/aa365247(v=vs.85).aspx

var reservedNames = []string{
	"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
	"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
	"CON", "PRN", "AUX", "NUL",
}

var Alt = AltMap{
	'<':  '＜', // Less-Than Sign -> Fullwidth Less-Than Sign
	'>':  '＞', // Greater-Than Sign -> Fullwidth Greater-Than Sign
	'*':  '＊', // Asterisk -> Fullwidth Asterisk
	'|':  '｜', // Vertical Line -> Fullwidth Vertical Line
	'/':  '／', // Solidus -> Fullwidth Solidus
	'\\': '＼', // Reverse Solidus -> Fullwidth Reverse Solidus
	'?':  '？', // Question Mark -> Fullwidth Question Mark
	':':  '：', // Colon -> Fullwidth Colon
	'"':  '＂', // Quotation Mark -> Fullwidth Quotation Mark
}

func (alt AltMap) Replace(s string) string {
	return strings.Map(func(r rune) rune {
		w := Alt[r]
		if illegal(w) {
			if illegal(r) {
				return -1
			}
			return r
		}
		return w
	}, s)
}

func illegal(r rune) bool {
	if 0 <= r && r <= 31 {
		return true
	}
	switch r {
	case '<', '>', '*', '|', '/', '\\', '?', ':', '"':
		return true
	}
	return false
}

func in(s string, set []string) bool {
	for _, t := range set {
		if strings.EqualFold(s, t) {
			return true
		}
	}
	return false
}

func split(name string) (string, string) {
	i := strings.LastIndexByte(name, '.')
	if i < 0 {
		return name, ""
	}
	return name[:i], name[i:]
}

func RegularName(name string) (string, bool) {
	zero := name
	name = strings.TrimSpace(name)
	name = Alt.Replace(name)
	for {
		base, ext := split(name)
		if in(base, reservedNames) {
			return zero, false
		}
		if ext != "." {
			break
		}
		name = name[:len(name)-len(".")]
		name = strings.TrimRightFunc(name, unicode.IsSpace)
	}
	return name, true
}

func IsRegular(name string) bool {
	raw := name
	name = strings.TrimSpace(name)
	if raw != name {
		return false
	}
	base, ext := split(name)
	if ext == "." {
		return false
	}
	if in(base, reservedNames) {
		return false
	}
	for _, r := range name {
		if illegal(r) {
			return false
		}
	}
	return true
}
