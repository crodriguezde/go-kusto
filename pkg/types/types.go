package types

type Column string

const (
	Bool     Column = "bool"
	DateTime Column = "datetime"
	Dynamic  Column = "dynamic"
	GUID     Column = "guid"
	Int      Column = "int"
	Long     Column = "long"
	Real     Column = "real"
	String   Column = "string"
	Timespan Column = "timespan"
	Decimal  Column = "decimal"
)

var valid = map[Column]bool{
	Bool:     true,
	DateTime: true,
	Dynamic:  true,
	GUID:     true,
	Int:      true,
	Long:     true,
	Real:     true,
	String:   true,
	Timespan: true,
	Decimal:  true,
}

func (c Column) IsValid() bool {
	return valid[c]
}
