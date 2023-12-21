package value

import "github.com/google/uuid"

type GUID struct {
	Value uuid.UUID

	Valid bool
}

func (GUID) isKustoVal() {}

func (g GUID) String() string {
	if !g.Valid {
		return ""
	}
	return g.Value.String()
}
