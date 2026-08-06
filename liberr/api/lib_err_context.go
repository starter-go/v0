package api

import "github.com/starter-go/i18n"

type Context struct {
	I18n i18n.Service

	Service Service

	Chain FilterChain

	FilterProviders []FilterRegistry

	DefaultLanguage i18n.Language
}
