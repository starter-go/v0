package subjects

import (
	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/localization"
)

type innerSetterImpl struct {

	// buffer  *Buffer

	context *Context
	agent   properties.Setter
}

func (inst *innerSetterImpl) init(c *Context) {
	if c == nil {
		panic("param: 'context' is nil")
	}
	inst.context = c
}

// SetAuthenticated implements Setter.
func (inst *innerSetterImpl) SetAuthenticated(authenticated bool) {

	const key = PNameAuthenticated
	agent := inst.innerGetAgent()
	agent.SetBool(string(key), authenticated)

}

// SetAvatar implements Setter.
func (inst *innerSetterImpl) SetAvatar(value string) {

	const key = PNameAvatar
	agent := inst.innerGetAgent()
	agent.SetString(string(key), value)

}

// SetDisplayName implements Setter.
func (inst *innerSetterImpl) SetDisplayName(name string) {

	const key = PNameNickName
	agent := inst.innerGetAgent()
	agent.SetString(string(key), name)

}

// SetLocale implements Setter.
func (inst *innerSetterImpl) SetLocale(l localization.Locale) {

	const key = PNameLanguage
	agent := inst.innerGetAgent()
	agent.SetString(string(key), string(l))

}

// SetNotAfter implements Setter.
func (inst *innerSetterImpl) SetNotAfter(t lang.Time) {

	const key = PNameNotAfter
	agent := inst.innerGetAgent()
	agent.SetInt64(string(key), t.Int())

}

// SetNotBefore implements Setter.
func (inst *innerSetterImpl) SetNotBefore(t lang.Time) {

	const key = PNameNotBefore
	agent := inst.innerGetAgent()
	agent.SetInt64(string(key), t.Int())

}

// SetProperty implements Setter.
func (inst *innerSetterImpl) SetProperty(name PropertyName, value string) {

	agent := inst.innerGetAgent()
	agent.SetString(string(name), value)

}

// SetSession implements Setter.
func (inst *innerSetterImpl) SetSession(se *rbac.SessionDTO) {

	// const key = PName
	// agent := inst.innerGetAgent()
	// agent.SetInt64(string(key), v)

	panic("no impl")

}

// SetUserEmail implements Setter.
func (inst *innerSetterImpl) SetUserEmail(addr string) {

	const key = PNameEmail
	agent := inst.innerGetAgent()
	agent.SetString(string(key), addr)

}

// SetUserID implements Setter.
func (inst *innerSetterImpl) SetUserID(uid rbac.UserID) {

	const key = PNameUserID
	agent := inst.innerGetAgent()
	agent.SetInt64(string(key), int64(uid))

}

// SetUserName implements Setter.
func (inst *innerSetterImpl) SetUserName(name rbac.UserName) {

	const key = PNameUserName
	agent := inst.innerGetAgent()
	agent.SetString(string(key), string(name))

}

func (inst *innerSetterImpl) innerGetAgent() properties.Setter {

	agent := inst.agent
	if agent != nil {
		return agent
	}

	ctx := inst.context
	buffer := ctx.Buffer
	if buffer == nil {
		buffer = new(Buffer)
	}

	pt := buffer.Properties
	if pt == nil {
		pt = properties.NewTable(nil)
	}

	buffer.Properties = pt
	agent = pt.Setter()
	ctx.Buffer = buffer

	inst.agent = agent
	return agent
}

func (inst *innerSetterImpl) _impl() Setter {
	return inst
}
