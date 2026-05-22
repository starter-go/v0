package subjects

import (
	"context"
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/localization"
)

type Getter interface {
	GetProperty(name string) string

	Names() []string

	GetSession(se *rbac.SessionDTO) bool

	IsAuthenticated() bool

	GetUserID() rbac.UserID

	GetUserName() rbac.UserName

	GetDisplayName() string

	GetUserEmail() string

	GetLocale() localization.Locale

	GetAvatar() string

	GetNotBefore() lang.Time

	GetNotAfter() lang.Time
}

type Setter interface {
	SetProperty(name, value string)

	SetSession(se *rbac.SessionDTO)

	SetAuthenticated(authenticated bool)

	SetUserID(uid rbac.UserID)

	SetUserName(name rbac.UserName)

	SetUserEmail(addr string)

	SetDisplayName(name string)

	SetLocale(l localization.Locale)

	SetAvatar(l string)

	SetNotBefore(t lang.Time)

	SetNotAfter(t lang.Time)
}

type Subject interface {
	GetContext() context.Context

	Reload() error

	Load() error

	Save() error

	Create() error

	Exit() error

	// // session
	// GetSession(result *rbac.SessionDTO) error
	// SetSession(se *rbac.SessionDTO) error
	// // properties
	// GetProperty(name string) (string, error)
	// SetProperty(name, value string) error
	// GetPropertyNameList() ([]string, error)

	DoGet() (Getter, error)

	DoSet() (Setter, error)
}

func GetCurrent(c context.Context) (Subject, error) {

	ada, err := GetAdapter(c)
	if err != nil {
		return nil, err
	}

	holder, err := ada.GetHolder(c)
	if err != nil {
		return nil, err
	}

	ctx := holder.Context
	if ctx == nil {
		msg := "subjects.GetCurrent() : context is nil"
		return nil, fmt.Errorf(msg)
	}

	sub := ctx.Facade
	if sub == nil {
		msg := "subjects.GetCurrent() : facade is nil"
		return nil, fmt.Errorf(msg)
	}

	return sub, nil
}
