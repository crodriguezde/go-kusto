package value

import "strconv"

type Real struct {
	Value float64

	Valid bool
}

func (Real) isKustoVal() {}

// String implements fmt.Stringer.
func (r Real) String() string {
	if !r.Valid {
		return ""
	}
	return strconv.FormatFloat(r.Value, 'e', -1, 64)
}
