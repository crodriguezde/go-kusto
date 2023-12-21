package value

type String struct {
	Value string
	Valid bool
}

func (String) isKustoVal() {}

func (s String) String() string {
	if !s.Valid {
		return ""
	}
	return s.Value
}
