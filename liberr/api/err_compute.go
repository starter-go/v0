package api

import "github.com/starter-go/base/lang"

func ComputeErrorURI(ns Namespace, name Name) lang.URI {
	s1 := string(ns)
	s2 := string(name)
	str := s1 + "#" + s2
	return lang.URI(str)
}
