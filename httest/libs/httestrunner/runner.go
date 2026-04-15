package httestrunner

import (
	"github.com/starter-go/v0/httest/libs/httestcore"
)

type Runner struct {
	context *Context
	rb      *httestcore.RequestBuilder
}

func (inst *Runner) Run() {

}

func (inst *Runner) innerTryInit() {

}

func (inst *Runner) SetMethod(method string) *Runner {

	inst.innerTryInit()

	inst.rb.SetMethod(method)

	return inst
}

func (inst *Runner) SetURL(url string) *Runner {

	inst.innerTryInit()

	inst.rb.SetURL(url)

	return inst
}

func (inst *Runner) SetBody(x string) *Runner {

	inst.innerTryInit()

	return inst
}

func (inst *Runner) SetHeader(name, value string) *Runner {

	inst.innerTryInit()

	return inst
}

func (inst *Runner) SetProperty(name, value string) *Runner {

	inst.innerTryInit()

	return inst
}
