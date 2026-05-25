package core

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/starter-go/v0/htttest"
	"github.com/starter-go/v0/htttest/api/h3t"
	"github.com/starter-go/vlog"
)

type UserAgentImpl struct {
	ac *h3t.AgentContext
}

// Execute implements h3t.UserAgent.
func (inst *UserAgentImpl) Execute(ctx *h3t.Transaction) error {

	defer func() {
		ctx.CloserPool.Close()
	}()

	request, err := inst.innerPrepareRequest(ctx)
	if err != nil {
		return inst.innerOnError(ctx, err)
	}

	client := inst.innerGetClient()
	inst.innerLogRequest(ctx, request)
	response, err := client.Do(request)
	inst.innerLogResponse(ctx, response)
	ctx.Have.Response = response
	if err != nil {
		return inst.innerOnError(ctx, err)
	}

	err = inst.innerHandleResponse(ctx, response)
	if err != nil {
		return inst.innerOnError(ctx, err)
	}

	err = inst.innerCheckResult(ctx)
	if err != nil {
		return inst.innerOnError(ctx, err)
	}

	return nil
}

// GetAgentContext implements h3t.UserAgent.
func (inst *UserAgentImpl) GetAgentContext() *h3t.AgentContext {
	return inst.ac
}

func (inst *UserAgentImpl) innerLogRequest(c *h3t.Transaction, req *http.Request) {

	if req == nil {
		return
	}

	method := req.Method
	location := req.URL
	vlog.Debug("HTTP %s %s", method, location)
}

func (inst *UserAgentImpl) innerLogResponse(c *h3t.Transaction, resp *http.Response) {

	if resp == nil {
		return
	}

	// code := resp.StatusCode
	msg := resp.Status
	vlog.Debug("  return   %s", msg)

}

func (inst *UserAgentImpl) innerPrepareRequest(c *h3t.Transaction) (*http.Request, error) {

	cc := c.CC
	method := c.Want.Method

	if method == "" {
		method = http.MethodGet
	}

	if cc == nil {
		cc = context.Background()
	}

	c.Want.Method = method
	c.CC = cc

	location, err := inst.innerPrepareRequestURL(c)
	if err != nil {
		return nil, err
	}
	c.Want.Location = location

	body, err := inst.innerPrepareRequestBody(c)
	c.CloserPool.Add(body)
	if err != nil {
		return nil, err
	}

	url := location.String()
	req, err := http.NewRequestWithContext(cc, method, url, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (inst *UserAgentImpl) innerPrepareRequestBody(c *h3t.Transaction) (io.ReadCloser, error) {

	body := c.Want.Body
	if body == nil {
		return nil, nil // no body
	}

	enc := body.Encoder
	if enc == nil {
		encoder, err := inst.innerFindEncoder(c, body.Type)
		if err != nil {
			return nil, err
		}
		enc = encoder
		body.Encoder = encoder
	}

	err := enc.Encode(body)
	if err != nil {
		return nil, err
	}

	bin := body.Encoded
	size := len(bin)
	body.Length = int64(size)

	buffer := bytes.NewBuffer(bin)
	rdr := io.NopCloser(buffer)
	return rdr, nil
}

func (inst *UserAgentImpl) innerPrepareRequestURL(c *h3t.Transaction) (*url.URL, error) {

	ac := c.AC
	resolver := ac.Resolver
	location, err := resolver.Resolve(c)
	if err != nil {
		return nil, err
	}

	c.Want.URL = location.String()
	c.Want.Location = location

	return location, nil
}

func (inst *UserAgentImpl) innerHandleResponse(c *h3t.Transaction, resp *http.Response) error {

	// status

	code := resp.StatusCode
	msg := resp.Status
	have := &c.Have

	// head

	headers := have.Head
	if headers == nil {
		headers = make(map[string]string)
	}

	for name, values := range resp.Header {
		name = strings.ToLower(name)
		for _, val := range values {
			headers[name] = val
		}
	}

	// body

	content, err := inst.innerHandleResponseBody(c, resp)
	if err != nil {
		return err
	}

	// keep
	have.Response = resp
	have.Status = code
	have.Message = msg
	have.Head = headers
	have.Body = content

	return nil
}

func (inst *UserAgentImpl) innerHandleResponseBody(c *h3t.Transaction, resp *http.Response) (*h3t.Content, error) {

	body := c.Have.Body
	if body == nil {
		body = new(h3t.Content)
	}

	ctypeStr := resp.Header.Get("content-type")

	body.Length = resp.ContentLength
	body.Type = h3t.ContentType(ctypeStr)

	src := resp.Body
	dst := bytes.NewBuffer(nil)
	maxLen := c.AC.MaxContentLength

	c.CloserPool.Add(src)

	if maxLen < body.Length {
		return nil, fmt.Errorf("[Error msg:'HTTP response body length is over-range'  max:%d have:%d]", maxLen, body.Length)
	}

	count, err := io.Copy(dst, src)
	if err != nil {
		return nil, err
	}

	if body.Length > 0 {
		if body.Length != count {
			return nil, fmt.Errorf("[Error msg:'bad HTTP response body length' want:%d have:%d]", body.Length, count)
		}
	}

	body.Encoded = dst.Bytes()

	// decode

	dec := body.Decoder
	if dec == nil {
		de2, err := inst.innerFindDecoder(c, body.Type)
		if err != nil {
			return nil, err
		}
		dec = de2
		body.Decoder = de2
	}
	if dec == nil {
		return nil, fmt.Errorf("[Error this:'UserAgentImpl' method:'innerHandleResponseBody' msg:'decoder is nil' ]")
	}

	err = dec.Decode(body)
	if err != nil {
		return nil, err
	}

	c.Have.Body = body
	return body, nil
}

func (inst *UserAgentImpl) innerFindDecoder(c *h3t.Transaction, ctype h3t.ContentType) (h3t.Decoder, error) {
	cm := c.Service.GetCodecManager()
	return cm.FindDecoder(ctype)
}

func (inst *UserAgentImpl) innerFindEncoder(c *h3t.Transaction, ctype h3t.ContentType) (h3t.Encoder, error) {
	cm := c.Service.GetCodecManager()
	return cm.FindEncoder(ctype)
}

func (inst *UserAgentImpl) innerOnError(c *h3t.Transaction, err error) error {
	if err != nil {
		c.Have.Error = err
	}
	return err
}

func (inst *UserAgentImpl) innerCheckResult(c *h3t.Transaction) error {
	code := c.Have.Status
	switch code {
	case http.StatusOK:
		return nil
	}
	msg := c.Have.Message
	return fmt.Errorf("HTTP " + msg)
}

func (inst *UserAgentImpl) innerGetClient() *http.Client {
	cli := inst.ac.Client
	if cli == nil {
		cli = http.DefaultClient
		inst.ac.Client = cli
	}
	return cli
}

func (inst *UserAgentImpl) _impl() htttest.UserAgent {
	return inst
}
