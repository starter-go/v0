package core

import (
	"context"
	"fmt"

	"github.com/starter-go/i18n"
	"github.com/starter-go/v0/liberr"
	"github.com/starter-go/v0/liberr/api"
)

type ErrorServiceImpl struct {

	//starter:component

	_as func(liberr.Service) //starter:as("#")

	Man     liberr.ErrorManager //starter:inject("#")
	I18nser i18n.Service        //starter:inject("#")
}

// GetManager implements [api.Service].
func (inst *ErrorServiceImpl) GetManager() api.ErrorManager {
	return inst.Man
}

// Error implements [api.Service].
func (inst *ErrorServiceImpl) Error(c context.Context, sel *api.Selector, args ...any) error {

	if sel == nil {
		return fmt.Errorf("liberr:Service:Error() : param 'Selector' is nil")
	}

	panic("unimplemented")

}

func (inst *ErrorServiceImpl) _impl() liberr.Service {
	return inst
}
