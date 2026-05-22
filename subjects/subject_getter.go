package subjects

import (
	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/localization"
)

type innerGetterImpl struct {
	context *Context
	cache   *Cache
	agent   properties.Getter
}

func (inst *innerGetterImpl) init(ctx *Context, cache *Cache) {

	if ctx == nil {
		panic("param: 'context' is nil")
	}

	if cache == nil {
		panic("param: 'cache' is nil")
	}

	inst.context = ctx
	inst.cache = cache

}

// GetAvatar implements Getter.
func (inst *innerGetterImpl) GetAvatar() string {

	const key = PNameAvatar
	agent := inst.innerGetAgent()
	return agent.GetString(string(key))
}

// GetDisplayName implements Getter.
func (inst *innerGetterImpl) GetDisplayName() string {

	const key = PNameNickName
	agent := inst.innerGetAgent()
	return agent.GetString(string(key))

}

// GetLocale implements Getter.
func (inst *innerGetterImpl) GetLocale() localization.Locale {

	const key = PNameAuthenticated
	agent := inst.innerGetAgent()
	value := agent.GetString(string(key))
	return localization.Locale(value)

}

// GetNotAfter implements Getter.
func (inst *innerGetterImpl) GetNotAfter() lang.Time {

	const key = PNameNotAfter
	return inst.innerGetTime(key)
}

// GetNotBefore implements Getter.
func (inst *innerGetterImpl) GetNotBefore() lang.Time {

	const key = PNameNotBefore
	return inst.innerGetTime(key)
}

// GetProperty implements Getter.
func (inst *innerGetterImpl) GetProperty(name string) string {

	agent := inst.innerGetAgent()
	return agent.GetString(name)
}

// GetSession implements Getter.
func (inst *innerGetterImpl) GetSession(se *rbac.SessionDTO) bool {

	const key = PNameAuthenticated
	agent := inst.innerGetAgent()
	// return
	//
	agent.GetBool(string(key))

	// panic("unimplemented")

	// todo : no impl

	return false
}

// GetUserEmail implements Getter.
func (inst *innerGetterImpl) GetUserEmail() string {

	const key = PNameEmail
	agent := inst.innerGetAgent()
	return agent.GetString(string(key))

}

// GetUserID implements Getter.
func (inst *innerGetterImpl) GetUserID() rbac.UserID {

	const key = PNameUserID
	agent := inst.innerGetAgent()
	value := agent.GetInt64(string(key))
	return rbac.UserID(value)

}

// GetUserName implements Getter.
func (inst *innerGetterImpl) GetUserName() rbac.UserName {

	const key = PNameUserName
	agent := inst.innerGetAgent()
	value := agent.GetString(string(key))
	return rbac.UserName(value)

}

// IsAuthenticated implements Getter.
func (inst *innerGetterImpl) IsAuthenticated() bool {

	const key = PNameAuthenticated
	agent := inst.innerGetAgent()
	return agent.GetBool(string(key))

}

// Names implements Getter.
func (inst *innerGetterImpl) Names() []string {

	agent := inst.innerGetAgent()
	return agent.ListItems("", "")

}

func (inst *innerGetterImpl) innerGetTime(name PropertyName) lang.Time {

	agent := inst.innerGetAgent()
	num := agent.GetInt64(string(name))
	return lang.Time(num)

}

func (inst *innerGetterImpl) innerGetAgent() properties.Getter {

	agent := inst.agent
	if agent != nil {
		return agent
	}

	cache := inst.cache
	pt := cache.Properties
	if pt == nil {
		pt = properties.NewTable(nil)
	}

	agent = pt.Getter()
	cache.Properties = pt
	inst.agent = agent

	return agent
}

func (inst *innerGetterImpl) _impl() Getter {
	return inst
}
