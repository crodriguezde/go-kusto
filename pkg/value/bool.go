package value

// Bool represents a Kusto boolean type. Bool implements Kusto.
type Bool struct {
	// Value holds the value of the type.
	Value bool
	// Valid indicates if this value was set.
	Valid bool
}

func (Bool) isKustoVal() {}

// String implements fmt.Stringer.
func (bo Bool) String() string {
	if !bo.Valid {
		return ""
	}
	if bo.Value {
		return "true"
	}
	return "false"
}
