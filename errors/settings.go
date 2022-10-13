package errors

type ErrorFormatter = func(err KError) string

type Settings struct {
	DefaultCode    int
	Separator      string
	ErrorFormatter ErrorFormatter
}

var CurrentSettings *Settings

func init() {
	CurrentSettings = new(Settings).ToDefault()
}

func (s *Settings) SetDefaultCode(code int) *Settings {
	s.DefaultCode = code
	return s
}

func (s *Settings) SetSeparator(separator string) *Settings {
	s.Separator = separator
	return s
}

func (s *Settings) SetErrorFormatter(fn ErrorFormatter) *Settings {
	s.ErrorFormatter = fn
	return s
}

func (s *Settings) ToDefault() *Settings {
	s.DefaultCode = 0
	s.Separator = ": "
	s.ErrorFormatter = nil
	return s
}
