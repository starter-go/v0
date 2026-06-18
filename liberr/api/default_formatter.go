package api

////////////////////////////////////////////////////////////////////////////////

var theDefaultFormatter Formatter

////////////////////////////////////////////////////////////////////////////////

func DefaultFormatter() Formatter {
	f := theDefaultFormatter
	if f == nil {
		f = GetBaseFormatter()
		theDefaultFormatter = f
	}
	return f
}

func SetDefaultFormatter(f Formatter) {
	theDefaultFormatter = f
}

func Format(parent error, reg *Registration, args ...any) error {
	f := DefaultFormatter()
	return f.Format(parent, reg, args...)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
