package subjects

import "context"

// the Input-Output-Context
type IOC struct {
	Context *Context

	CC context.Context

	Want Want

	Have Have
}
