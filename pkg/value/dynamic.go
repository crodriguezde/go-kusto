package value

type Dynamic struct {
	Value []byte

	Valid bool
}

func (Dynamic) isKustoVal() {}

func (d Dynamic) String() string {
	if !d.Valid {
		return ""
	}

	return string(d.Value)
}
