package subjects

import "net/http"

type Method string

const (
	MethodGet    Method = http.MethodGet
	MethodPut    Method = http.MethodPut
	MethodPost   Method = http.MethodPost
	MethodDelete Method = http.MethodDelete
	MethodReload Method = "RELOAD"
)
