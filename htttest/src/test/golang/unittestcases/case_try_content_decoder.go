package unittestcases

import (
	"net/http"

	"github.com/starter-go/units"
	"github.com/starter-go/v0/htttest"
	"github.com/starter-go/vlog"
)

type CaseTryContentDecoder struct {

	//starter:component

	Service htttest.Service //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *CaseTryContentDecoder) ListRegistrations(list []*units.Registration) []*units.Registration {

	const id = "case-try-content-decoder"

	r1 := &units.Registration{

		Name:     id + "-json",
		Enabled:  true,
		Priority: 1,

		Do: inst.runAsJson,
	}

	r2 := &units.Registration{

		Name:     id + "-text",
		Enabled:  true,
		Priority: 2,

		Do: inst.runAsText,
	}

	list = append(list, r1, r2)

	return list
}

func (inst *CaseTryContentDecoder) runAsText() error {

	ser := inst.Service
	agent := ser.GetUserAgent()
	tr := ser.NewTransaction(http.MethodGet, "/api/v1/auth")
	decoder := inst.findDecoder("text/plain")

	data1 := new(InnerAuthVO)
	data2 := new(InnerAuthVO)
	tr.Want.Body = &htttest.Content{
		Type: "application/json",
		Data: data1,
	}
	tr.Have.Body = &htttest.Content{
		Data:    data2,
		Decoder: decoder,
	}

	defer func() {
		inst.handleResponseData(tr)
		data := tr.Have.Body.Data
		str, ok := data.(string)
		if ok {
			vlog.Debug("Response Data (as string) : %s", str)
		}
	}()

	return agent.Execute(tr)
}

func (inst *CaseTryContentDecoder) runAsJson() error {

	ser := inst.Service
	agent := ser.GetUserAgent()
	tr := ser.NewTransaction(http.MethodGet, "/api/v1/auth")
	decoder := inst.findDecoder("application/json")

	data1 := new(InnerAuthVO)
	data2 := new(InnerAuthVO)
	tr.Want.Body = &htttest.Content{
		Type: "application/json",
		Data: data1,
	}
	tr.Have.Body = &htttest.Content{
		Data:    data2,
		Decoder: decoder,
	}

	defer func() {
		inst.handleResponseData(tr)
	}()

	return agent.Execute(tr)
}

func (inst *CaseTryContentDecoder) findDecoder(ctype htttest.ContentType) htttest.Decoder {
	cm := inst.Service.GetCodecManager()
	dec, _ := cm.FindDecoder(ctype)
	return dec
}

func (inst *CaseTryContentDecoder) handleResponseData(tr *htttest.Transaction) error {
	return nil
}

func (inst *CaseTryContentDecoder) _impl() units.Unit {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

// type InnerAuthVO struct {
// 	Foo int
// 	Bar string
// }

////////////////////////////////////////////////////////////////////////////////
