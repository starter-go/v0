package httestcore

import "net/http"

type RequestBuilder struct {
	method  string
	url     string
	headers map[string]string
}

func (inst *RequestBuilder) SetMethod(method string) *RequestBuilder {

	inst.method = method

	return inst
}

func (inst *RequestBuilder) SetURL(url string) *RequestBuilder {

	inst.url = url

	return inst
}

func (inst *RequestBuilder) SetHeader(name, value string) *RequestBuilder {

	t := inst.innerPrepareStringMap(inst.headers)
	t[name] = value
	inst.headers = t

	return inst
}

func (inst *RequestBuilder) SetBody(body *string) *RequestBuilder {

	return inst

}

func (inst *RequestBuilder) Build() *http.Request {

	req := new(http.Request)

	return req
}

func (inst *RequestBuilder) innerPrepareStringMap(t map[string]string) map[string]string {
	if t == nil {
		t = make(map[string]string)
	}
	return t
}
