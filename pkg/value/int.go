package value

import "strconv"

type Int struct {
	Value int32
	Valid bool
}

func (Int) isKustoVal() {}

func (in Int) String() string {
	if !in.Valid {
		return ""
	}
	return strconv.Itoa(int(in.Value))
}
