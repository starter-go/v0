package httestrunner

import (
	"fmt"
	"net/http"
)

////////////////////////////////////////////////////////////////////////////////

type runnerInner struct {
	context *Context
}

func (inst *runnerInner) init(c *Context) {
	inst.context = c
}

func (inst *runnerInner) run() error {

	ctx := inst.context
	rb := ctx.RequestBuilder

	req := rb.Build()
	client := http.DefaultClient

	inst.innerLogRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	inst.innerLogResponse(resp)

	return inst.innerCheckResponse(resp)
}

func (inst *runnerInner) innerLogRequest(req *http.Request) {}

func (inst *runnerInner) innerLogResponse(resp *http.Response) {}

func (inst *runnerInner) innerCheckResponse(resp *http.Response) error {

	code := resp.StatusCode

	if code == http.StatusOK {
		return nil
	}

	text := resp.Status
	return fmt.Errorf("HTTP %s", text)
}

////////////////////////////////////////////////////////////////////////////////
