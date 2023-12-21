package value

import (
	"fmt"
	"time"
)

type DateTime struct {
	Value time.Time

	Valid bool
}

func (d DateTime) String() string {
	if !d.Valid {
		return ""
	}
	return fmt.Sprint(d.Value.Format(time.RFC3339Nano))
}

func (DateTime) isKustoVal() {}
