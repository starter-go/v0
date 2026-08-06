package api

import "fmt"

////////////////////////////////////////////////////////////////////////////////

var theBaseFormatter innerBaseFormatter

////////////////////////////////////////////////////////////////////////////////

func GetBaseFormatter() Formatter {
	return &theBaseFormatter
}

////////////////////////////////////////////////////////////////////////////////

type innerBaseFormatter struct {
}

// Format implements [Formatter].
func (inst *innerBaseFormatter) Format(args ...any) error {
	panic("unimplemented")
}

// FormatInfo implements [Formatter].
func (inst *innerBaseFormatter) FormatInfo(info *HyperErrorInfo, format string, args ...any) string {
	head := ""
	if info == nil {
		head = "HyperError"
	} else {
		name := info.Name
		code := info.Code
		head = fmt.Sprintf("%s(%d)", name, code)
	}
	body := fmt.Sprintf(format, args...)
	return (head + ":" + body)
}

// FormatRegistration implements [Formatter].
func (inst *innerBaseFormatter) FormatRegistration(reg *Registration, args ...any) string {
	if reg == nil {
		return ""
	}
	a1i := &reg.HyperErrorInfo
	a2f := reg.Format
	return inst.FormatInfo(a1i, a2f, args...)
}

func (inst *innerBaseFormatter) _impl() Formatter {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
