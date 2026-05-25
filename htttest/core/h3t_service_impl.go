package core

import (
	"github.com/starter-go/application/attributes"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/v0/htttest"
	"github.com/starter-go/v0/htttest/api/h3t"
)

type ServiceImpl struct {

	//starter:component

	_as func(htttest.Service) //starter:as("#")

	BaseURL          string //starter:inject("${starter.htttest.base-url}")
	MaxContentLength int    //starter:inject("${starter.htttest.max-content-length}")

	CodecManager htttest.CodecManager //starter:inject("#")

	ac *h3t.AgentContext
}

// GetCodecManager implements h3t.Service.
func (inst *ServiceImpl) GetCodecManager() h3t.CodecManager {
	return inst.CodecManager
}

// GetAgentContext implements h3t.Service.
func (inst *ServiceImpl) GetAgentContext() *h3t.AgentContext {
	ac := inst.innerGetAgentContext(true)
	return ac
}

// GetUserAgent implements h3t.Service.
func (inst *ServiceImpl) GetUserAgent() h3t.UserAgent {
	ac := inst.innerGetAgentContext(true)
	return ac.UserAgent
}

func (inst *ServiceImpl) innerGetAgentContext(required bool) *h3t.AgentContext {
	ac, err := inst.innerGetAC()
	if err != nil {
		if required {
			panic(err)
		} else {
			return nil
		}
	}
	return ac
}

func (inst *ServiceImpl) innerGetAC() (*h3t.AgentContext, error) {
	ac := inst.ac
	if ac == nil {
		ac2, err := inst.innerLoadAC()
		if err != nil {
			return nil, err
		}
		ac = ac2
		inst.ac = ac2
	}
	return ac, nil
}

func (inst *ServiceImpl) innerLoadAC() (*h3t.AgentContext, error) {

	ac := new(h3t.AgentContext)
	agent := new(UserAgentImpl)
	resolver := new(DefaultUrlResolver)

	agent.ac = ac

	ac.Resolver = resolver
	ac.Service = inst
	ac.UserAgent = agent
	ac.BaseURL = inst.BaseURL
	ac.MaxContentLength = int64(inst.MaxContentLength)
	ac.Attributes = attributes.NewTable(nil)
	ac.Properties = properties.NewTable(nil)
	ac.RequestHeaders = properties.NewTable(nil)

	return ac, nil
}

// NewRequest implements h3t.Service.
func (inst *ServiceImpl) NewTransaction(method string, url string) *h3t.Transaction {

	ac := inst.ac
	tr := new(h3t.Transaction)
	want := &tr.Want
	have := &tr.Have

	tr.AC = ac
	tr.Service = inst

	want.Method = method
	want.URL = url
	want.Head = make(map[string]string)
	want.Params = make(map[string]string)
	want.Query = make(map[string]string)

	have.Head = make(map[string]string)

	return tr
}

func (inst *ServiceImpl) init() {
	inst.innerGetAC()
}

func (inst *ServiceImpl) _impl() htttest.Service {
	return inst
}
