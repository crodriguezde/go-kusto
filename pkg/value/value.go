// Package value provides an interface and a type for handling Kusto values.
package value

import "reflect"

// Kusto is an interface that represents a Kusto value.
// It provides methods for checking if a value is a Kusto value,
// converting it to a string, and converting it from a reflect.Value.
type Value interface {
	isKustoVal()                   // Checks if the value is a Kusto value
	String() string                // Converts the Kusto value to a string
	Convert(v reflect.Value) error // Converts a reflect.Value to a Kusto value
}

// Values is a slice of Kusto values.
type Values []Value
