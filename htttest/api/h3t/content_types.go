package h3t

import "strings"

type ContentType string

func (ct ContentType) String() string {
	return string(ct)
}

func (ct ContentType) Pure() ContentType {
	const (
		sep = ";"
		n   = 2
	)
	str := ct.String()
	if !strings.ContainsAny(str, sep) {
		return ct
	}
	elist := strings.SplitN(str, sep, n)
	size := len(elist)
	if size == 1 || size == 2 {
		str = elist[0]
	}
	str = strings.TrimSpace(str)
	return ContentType(str)
}

func (ct ContentType) Normalize() ContentType {
	str := ct.String()
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	return ContentType(str)
}
