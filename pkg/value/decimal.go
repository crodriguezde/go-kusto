package value

import "regexp"

var DecRE = regexp.MustCompile(`^((\d+\.?\d*)|(\d*\.?\d+))$`) // Matches decimal numbers, with or without decimal dot, with optional parts missing.

type Decimal struct {
	Value string
	Valid bool
}

func (Decimal) isKustoVal() {}

func (d Decimal) String() string {
	if !d.Valid {
		return ""
	}
	return d.Value
}
