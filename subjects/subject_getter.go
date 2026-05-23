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
func (inst *innerGetterImpl) GetProperty(name PropertyName) string {

	agent := inst.innerGetAgent()
	return agent.GetString(string(name))
}

// GetSession implements Getter.
func (inst *innerGetterImpl) GetSession(se *rbac.SessionDTO) bool {

	agent := inst.innerGetAgent()

	// keys

	const (
		keyNotAfter  = string(PNameNotAfter)
		keyNotBefore = string(PNameNotBefore)
		keyCreatedAt = string(PNameCreatedAt)
		keyUpdatedAt = string(PNameUpdatedAt)

		keySessionID     = string(PNameSessionID)
		keySessionUUID   = string(PNameSessionUUID)
		keyUserID        = string(PNameUserID)
		keyUserName      = string(PNameUserName)
		keyNickName      = string(PNameNickName)
		keyAvatar        = string(PNameAvatar)
		keyRoles         = string(PNameRoles)
		keyLanguage      = string(PNameLanguage)
		keyEmail         = string(PNameEmail)
		keyAuthenticated = string(PNameAuthenticated)
	)

	// get

	valueAuthenticated := agent.GetBool(keyAuthenticated)
	valueNotAfter := agent.GetInt64(keyNotAfter)
	valueNotBefore := agent.GetInt64(keyNotBefore)
	valueCreatedAt := agent.GetInt64(keyCreatedAt)
	valueUpdatedAt := agent.GetInt64(keyUpdatedAt)
	valueSessionID := agent.GetInt64(keySessionID)
	valueSessionUUID := agent.GetString(keySessionUUID)
	valueUserName := agent.GetString(keyUserName)
	valueUserID := agent.GetInt64(keyUserID)
	valueNickName := agent.GetString(keyNickName)
	valueAvatar := agent.GetString(keyAvatar)
	valueRoles := agent.GetString(keyRoles)
	valueEmail := agent.GetString(keyEmail)
	valueLang := agent.GetString(keyLanguage)

	// set

	se.Authenticated = valueAuthenticated
	se.CreatedAt = lang.Time(valueCreatedAt)
	se.UpdatedAt = lang.Time(valueUpdatedAt)
	se.StartedAt = lang.Time(valueNotBefore)
	se.ExpiredAt = lang.Time(valueNotAfter)
	se.UUID = lang.UUID(valueSessionUUID)
	se.ID = rbac.SessionID(valueSessionID)
	se.UserID = rbac.UserID(valueUserID)
	se.Owner = rbac.UserID(valueUserID)
	se.Username = rbac.UserName(valueUserName)
	se.Nickname = valueNickName
	se.Avatar = valueAvatar
	se.Email = rbac.EmailAddress(valueEmail)
	se.Roles = rbac.RoleNameList(valueRoles)
	se.Language = localization.Locale(valueLang)

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
