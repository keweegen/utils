package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultSettings(t *testing.T) {
	s := *CurrentSettings.ToDefault()

	assert.Equal(t, 0, s.DefaultCode, "default code")
	assert.Equal(t, ": ", s.Separator, "separator")
	assert.Nil(t, s.ErrorFormatter, "error formatter")
	assert.EqualValues(t, &s, CurrentSettings)
}

func TestCustomSettings(t *testing.T) {
	s := *CurrentSettings.ToDefault().
		SetDefaultCode(200).
		SetSeparator(" -> ").
		SetErrorFormatter(nil)

	assert.Equal(t, 200, s.DefaultCode, "default code")
	assert.Equal(t, " -> ", s.Separator, "separator")
	assert.Nil(t, s.ErrorFormatter, "error formatter")
	assert.EqualValues(t, &s, CurrentSettings)

	CurrentSettings.ToDefault()

	assert.NotEqualValues(t, &s, CurrentSettings)
}
