package errors

import (
    "errors"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNewKErrorWithDefaultSettings(t *testing.T) {
    err := New(10, "myerror")
    errTemplate := "[ERR%d] %s"
    errTemplateWithParent := "[ERR%d] %s: %s"

    cases := []struct {
        name                     string
        err                      error
        expectedMessage          string
        expectedUnwrap           error
        expectedUnwrappedMessage string
    }{
        {
            name:                     "without parent",
            err:                      err,
            expectedMessage:          fmt.Sprintf(errTemplate, 10, "myerror"),
            expectedUnwrap:           nil,
            expectedUnwrappedMessage: "",
        },
        {
            name:                     "with parent",
            err:                      New(20, "withparent", err),
            expectedMessage:          fmt.Sprintf(errTemplateWithParent, 20, "withparent", err.Error()),
            expectedUnwrap:           err,
            expectedUnwrappedMessage: err.Error(),
        },
        {
            name:                     "with native error",
            err:                      fmt.Errorf("native error: %w", err),
            expectedMessage:          fmt.Sprintf("native error: %s", err.Error()),
            expectedUnwrap:           err,
            expectedUnwrappedMessage: err.Error(),
        },
        {
            name:                     "with parent native error",
            err:                      Wrap(errors.New("native error"), 30, "withnativeparent"),
            expectedMessage:          fmt.Sprintf(errTemplateWithParent, 30, "withnativeparent", "native error"),
            expectedUnwrap:           errors.New("native error"),
            expectedUnwrappedMessage: "native error",
        },
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            assert.Equal(t, c.expectedMessage, c.err.Error(), "message")
            assert.Equal(t, c.expectedUnwrap, errors.Unwrap(c.err), "unwrap")
        })
    }
}

func TestNewKErrorWithCustomSettings(t *testing.T) {
    s := CurrentSettings.ToDefault().
        SetDefaultCode(-1).
        SetSeparator(" -> ").
        SetErrorFormatter(func(err KError) string {
            return fmt.Sprintf("%d", err.code)
        })

    err := New(10, "myerror")
    errTemplate := "%d"
    errTemplateWithParent := "%d -> %s"

    cases := []struct {
        name                     string
        err                      error
        expectedMessage          string
        expectedUnwrap           error
        expectedUnwrappedMessage string
    }{
        {
            name:                     "without parent",
            err:                      err,
            expectedMessage:          fmt.Sprintf(errTemplate, 10),
            expectedUnwrap:           nil,
            expectedUnwrappedMessage: "",
        },
        {
            name:                     "with parent",
            err:                      New(20, "withparent", err),
            expectedMessage:          fmt.Sprintf(errTemplateWithParent, 20, err.Error()),
            expectedUnwrap:           err,
            expectedUnwrappedMessage: err.Error(),
        },
        {
            name:                     "with zero value code",
            err:                      New(0, "zerovalue"),
            expectedMessage:          fmt.Sprintf(errTemplate, s.DefaultCode),
            expectedUnwrap:           nil,
            expectedUnwrappedMessage: "",
        },
        {
            name:                     "with native error",
            err:                      fmt.Errorf("native error: %w", err),
            expectedMessage:          fmt.Sprintf("native error: %s", err.Error()),
            expectedUnwrap:           err,
            expectedUnwrappedMessage: err.Error(),
        },
        {
            name:                     "with parent native error",
            err:                      Wrap(errors.New("native error"), 30, "withnativeparent"),
            expectedMessage:          fmt.Sprintf(errTemplateWithParent, 30, "native error"),
            expectedUnwrap:           errors.New("native error"),
            expectedUnwrappedMessage: "native error",
        },
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            assert.Equal(t, c.expectedMessage, c.err.Error(), "message")
            assert.Equal(t, c.expectedUnwrap, errors.Unwrap(c.err), "unwrap")

            if c.expectedUnwrap != nil {
                assert.Equal(t, c.expectedUnwrappedMessage, errors.Unwrap(c.err).Error())
            }
        })
    }
}
