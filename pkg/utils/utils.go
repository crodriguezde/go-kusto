package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func QuoteString(value string, hidden bool) string {
	if value == "" {
		return value
	}

	var literal strings.Builder

	if hidden {
		literal.WriteString("h")
	}
	literal.WriteString("\"")
	for _, c := range value {
		switch c {
		case '\'':
			literal.WriteString("\\'")

		case '"':
			literal.WriteString("\\\"")

		case '\\':
			literal.WriteString("\\\\")

		case '\x00':
			literal.WriteString("\\0")

		case '\a':
			literal.WriteString("\\a")

		case '\b':
			literal.WriteString("\\b")

		case '\f':
			literal.WriteString("\\f")

		case '\n':
			literal.WriteString("\\n")

		case '\r':
			literal.WriteString("\\r")

		case '\t':
			literal.WriteString("\\t")

		case '\v':
			literal.WriteString("\\v")

		default:
			if !ShouldBeEscaped(c) {
				literal.WriteString(string(c))
			} else {
				literal.WriteString(fmt.Sprintf("\\u%04x", c))
			}

		}
	}
	literal.WriteString("\"")

	return literal.String()
}

// ShouldBeEscaped Checks whether a rune should be escaped or not based on it's type.
func ShouldBeEscaped(c int32) bool {
	if c <= unicode.MaxLatin1 {
		return unicode.IsControl(c)
	}
	return true
}
