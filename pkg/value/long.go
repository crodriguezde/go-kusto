package value

import "strconv"

type Long struct {
	Value int64

	Valid bool
}

func (Long) isKustoVal() {}

func (l Long) String() string {
	if !l.Valid {
		return ""
	}
	return strconv.Itoa(int(l.Value))
}
