package subjects

type FilterPriority int

const (
	FilterPriorityMin FilterPriority = iota

	FilterPriorityDB
	FilterPriorityCache2 // with redis
	FilterPriorityJWT
	FilterPriorityBuffer
	FilterPriorityCodec
	FilterPriorityGoodResult
	FilterPriorityLog
	FilterPriorityCache1 // with runtime memory

	FilterPriorityMax
)

////////////////////////////////////////////////////////////////////////////////

type FilterRegistration struct {
	Name string

	Enabled bool

	Priority FilterPriority

	Writer WriteFilter

	Reader ReadFilter
}

type FilterRegistry interface {
	GetRegistrationList() []*FilterRegistration
}
