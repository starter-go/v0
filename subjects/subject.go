package subjects

import (
	"context"
	"fmt"

	"github.com/starter-go/application/properties"
)

type Subject interface {
	GetProperties() (properties.Table, error)

	Reload() error

	Create() error

	Save() error
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
