package core

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/starter-go/v0/htttest"
	"github.com/starter-go/v0/htttest/api/h3t"
)

type DefaultUrlResolver struct {
}

// Resolve implements h3t.Resolver.
func (inst *DefaultUrlResolver) Resolve(c *h3t.Transaction) (*url.URL, error) {

	resolving := new(innerResolving)
	err := resolving.resolve(c)
	if err != nil {
		return nil, err
	}
	return resolving.getResult()
}

func (inst *DefaultUrlResolver) _impl() htttest.Resolver {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerResolving struct {
	raw   string
	base1 string
	tmp1  string

	result *url.URL
	tmp    *url.URL
	base   *url.URL
	params map[string]string
	query  map[string]string
}

func (inst *innerResolving) init(c *h3t.Transaction) error {

	want := &c.Want
	ac := c.AC

	u1 := want.RawURL
	u2 := want.URL
	u3 := want.Location

	// var err error

	if u3 != nil {
		str := u3.String()
		u1 = str
		u2 = str
	} else if u1 == "" {
		u1 = u2
	} else if u2 == "" {
		u2 = u1
	}

	inst.raw = u1
	inst.tmp = nil
	inst.base = nil
	inst.result = nil
	inst.base1 = ac.BaseURL
	inst.query = c.Want.Query
	inst.params = c.Want.Params

	if inst.base1 == "" {
		return fmt.Errorf("base URL is empty")
	}
	if inst.raw == "" {
		return fmt.Errorf("raw URL is empty")
	}

	return nil
}

func (inst *innerResolving) resolve(c *h3t.Transaction) error {

	err := inst.init(c)
	if err != nil {
		return err
	}

	steps := make([]func() error, 0)

	steps = append(steps, inst.stepDoInjectParams)
	steps = append(steps, inst.stepDoReference)
	steps = append(steps, inst.stepDoInjectQuery)
	steps = append(steps, inst.stepDoMakeResult)

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
		if inst.hasResult() {
			return nil
		}
	}
	return nil
}

func (inst *innerResolving) stepDoInjectQuery() error {
	src := inst.query
	if src == nil {
		return nil
	}
	if len(src) == 0 {
		return nil
	}
	tmp := inst.tmp.Query()
	for k, v := range src {
		tmp.Set(k, v)
	}
	inst.tmp.RawQuery = tmp.Encode()
	return nil
}

func (inst *innerResolving) stepDoInjectParams() error {

	const (
		markBegin = "{{"
		markEnd   = "}}"
	)

	raw := inst.raw // raw-string
	parts := strings.Split(raw, markBegin)
	builder := new(strings.Builder)

	for index, part := range parts {
		if index == 0 {
			if strings.Contains(part, markEnd) {
				return fmt.Errorf("bad URL format to resolve '%s'", raw)
			}
			builder.WriteString(part)
		} else {
			p0p1 := strings.Split(part, markEnd)
			p0 := p0p1[0]
			p1 := p0p1[1]
			value := inst.innerGetParam(p0)
			if value == "" {
				return fmt.Errorf("no required param '%s' to resolve URL  '%s' ", p0, raw)
			}
			builder.WriteString(value)
			builder.WriteString(p1)
		}
	}

	str2 := builder.String()
	u2, err := url.Parse(str2)
	if err != nil {
		return err
	}
	inst.tmp = u2
	inst.tmp1 = str2
	return nil
}

func (inst *innerResolving) hasResult() bool {
	return inst.result != nil
}

func (inst *innerResolving) getResult() (*url.URL, error) {
	res := inst.result
	if res == nil {
		return nil, fmt.Errorf("DefaultUrlResolver: result URL is nil")
	}
	return res, nil
}

func (inst *innerResolving) stepDoReference() error {

	base := inst.base

	if base == nil {
		str := inst.base1
		b2, err := url.Parse(str)
		if err != nil {
			return err
		}
		base = b2
		inst.base = base
	}

	ref := inst.tmp
	u2 := base.ResolveReference(ref)

	inst.tmp = u2
	return nil
}

func (inst *innerResolving) innerGetParam(name string) string {
	table := inst.params
	if table == nil {
		return ""
	}
	name = strings.TrimSpace(name)
	return table[name]
}

func (inst *innerResolving) stepDoMakeResult() error {
	inst.result = inst.tmp
	return nil
}

////////////////////////////////////////////////////////////////////////////////
