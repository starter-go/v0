package statestores

type FilterRegistry interface {
	ListRegistrations() []*FilterRegistration
}

type FilterRegistration struct {
	Name string

	Enabled bool

	Priority FilterPriority

	Filter Filter
}
