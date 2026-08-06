package api

import (
	"github.com/starter-go/i18n"
)

type Selector struct {
	NS Namespace

	Name Name

	Lang i18n.Language
}

type Service interface {
	Error(w *Want) error

	GetManager() ErrorManager
}
