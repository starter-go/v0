package api

// the code of error
type Code int32

// the name of error
type Name string

// the ns of error
type Namespace string

type Argument struct {
	Name        string
	Value       any
	Description string
}

type Registration struct {

	// core-info:
	HyperErrorInfo

	Format     string
	FormatI18n string      // a i18n-string-id to query format string
	Args       []*Argument // items like: "['id(int)','name(string)','sum([]byte)', ... ]"
	Formatter  Formatter
	Example    error
}

type ErrorSet interface {
	ListErrors() []*Registration
}

type Formatter interface {
	Format(args ...any) error
}

////////////////////////////////////////////////////////////////////////////////

func (inst *Registration) NewInfo() *HyperErrorInfo {
	info := new(HyperErrorInfo)
	*info = inst.HyperErrorInfo
	return info
}

func (inst *Registration) GetFormatter() Formatter {

	return nil
}

////////////////////////////////////////////////////////////////////////////////
