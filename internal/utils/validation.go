package utils

import (
	"net/url"
	"strings"
	"unicode"

	"github.com/rl404/fairy/validation"
	"github.com/rl404/fairy/validation/playground"
	"github.com/rl404/mal-cover/internal/errors"
)

var val validation.Validator

func init() {
	val = playground.New(true)
	val.RegisterModifier("no_space", modNoSpace)
	val.RegisterModifier("unescape", modUnescape)
	val.RegisterValidator("style", valStyle)
	val.RegisterValidatorError("required", valErrRequired)
	val.RegisterValidatorError("gt", valErrGT)
	val.RegisterValidatorError("oneof", valErrOneOf)
	val.RegisterValidatorError("style", valErrStyle)
}

// Validate to validate struct using validate tag.
// Use pointer.
func Validate(data interface{}) error {
	return val.Validate(data)
}

func modNoSpace(in string, _ ...string) string {
	return strings.ReplaceAll(in, " ", "")
}

func modUnescape(in string, _ ...string) string {
	in, _ = url.QueryUnescape(in)
	return in
}

func valStyle(f interface{}, param ...string) bool {
	return f.(string) != ""
}

func valErrRequired(f string, param ...string) error {
	return errors.ErrRequiredField(camelToSnake(f))
}

func valErrGT(f string, param ...string) error {
	return errors.ErrGTField(camelToSnake(f), param[0])
}

func valErrOneOf(f string, param ...string) error {
	return errors.ErrOneOfField(camelToSnake(f), param[0])
}

func valErrStyle(f string, param ...string) error {
	return errors.ErrStyleField()
}

func camelToSnake(name string) string {
	if name == "" {
		return ""
	}

	runes := []rune(name)
	var out strings.Builder

	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 {
				prev := runes[i-1]
				next := rune(0)
				if i+1 < len(runes) {
					next = runes[i+1]
				}

				if unicode.IsLower(prev) ||
					(unicode.IsUpper(prev) && unicode.IsLower(next)) {
					out.WriteByte('_')
				}
			}

			out.WriteRune(unicode.ToLower(r))
			continue
		}

		out.WriteRune(r)
	}

	return out.String()
}
