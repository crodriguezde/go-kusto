package query

import (
	"fmt"
	"math/big"
	"time"

	"github.com/crodriguezde/go-kusto/pkg/errors"
	"github.com/crodriguezde/go-kusto/pkg/types"
	"github.com/crodriguezde/go-kusto/pkg/utils"
	"github.com/crodriguezde/go-kusto/pkg/value"
	"github.com/google/uuid"
)

type ParamType struct {
	Type    types.Column
	Default interface{}
	name    string
}

type ParamTypes map[string]ParamType

func (p ParamType) validate() error {
	if !p.Type.IsValid() {
		return errors.ErrInvalidType
	}

	if p.Default == nil {
		return nil
	}

	switch p.Type {
	case types.Bool:
		if _, ok := p.Default.(bool); !ok {
			return errors.ErrWrapf(errors.ErrInvalidType, "expected %s, got %T", p.Type, p.Default)
		}
		return nil
	case types.DateTime:
		if _, ok := p.Default.(time.Time); !ok {
			return errors.ErrWrapf(errors.ErrInvalidType, "expected %s, got %T", p.Type, p.Default)
		}
		return nil
	case types.Dynamic:
		return errors.ErrWrapf(errors.ErrInvalidType, "cannot set default value for dynamic type")
	case types.GUID:
		if _, ok := p.Default.(uuid.UUID); !ok {
			return errors.ErrWrapf(errors.ErrInvalidType, "expected %s, got %T", p.Type, p.Default)
		}
		return nil
	case types.Int:
		if _, ok := p.Default.(int32); !ok {
			return errors.ErrWrapf(errors.ErrInvalidType, "expected %s, got %T", p.Type, p.Default)
		}
		return nil
	case types.Long:
		if _, ok := p.Default.(int64); !ok {
			return errors.ErrWrapf(errors.ErrInvalidType, "expected %s, got %T", p.Type, p.Default)
		}
		return nil
	case types.Real:
		if _, ok := p.Default.(float64); !ok {
			return errors.ErrWrapf(errors.ErrInvalidType, "expected %s, got %T", p.Type, p.Default)
		}
		return nil
	case types.String:
		if _, ok := p.Default.(string); !ok {
			return errors.ErrWrapf(errors.ErrInvalidType, "expected %s, got %T", p.Type, p.Default)
		}
		return nil
	case types.Timespan:
		if _, ok := p.Default.(time.Duration); !ok {
			return errors.ErrWrapf(errors.ErrInvalidType, "expected %s, got %T", p.Type, p.Default)
		}
		return nil
	case types.Decimal:
		switch v := p.Default.(type) {
		case string:
			if !value.DecRE.MatchString(v) {
				return errors.ErrWrapf(errors.ErrInvalidType, "string representing decimal does not appear to be a decimal number, was %v", v)
			}
			return nil
		case *big.Float:
			if v == nil {
				return errors.ErrWrapf(errors.ErrInvalidType, "*big.Float type cannot be set to the nil value")
			}
			return nil
		case *big.Int:
			if v == nil {
				return errors.ErrWrapf(errors.ErrInvalidType, "*big.Int type cannot be set to the nil value")
			}
			return nil
		}
	}

	return errors.ErrWrapf(errors.ErrInvalidType, "unknown type %s", p.Type)
}

func (p ParamType) string() string {
	switch p.Type {
	case types.Bool:
		if p.Default == nil {
			return p.name + ":bool"
		}
		v := p.Default.(bool)
		return fmt.Sprintf("%s:bool = bool(%v)", p.name, v)
	case types.DateTime:
		if p.Default == nil {
			return p.name + ":datetime"
		}
		v := p.Default.(time.Time)
		return fmt.Sprintf("%s:datetime = datetime(%s)", p.name, v.Format(time.RFC3339Nano))
	case types.Dynamic:
		return p.name + ":dynamic"
	case types.GUID:
		if p.Default == nil {
			return p.name + ":guid"
		}
		v := p.Default.(uuid.UUID)
		return fmt.Sprintf("%s:guid = guid(%s)", p.name, v.String())
	case types.Int:
		if p.Default == nil {
			return p.name + ":int"
		}
		v := p.Default.(int32)
		return fmt.Sprintf("%s:int = int(%d)", p.name, v)
	case types.Long:
		if p.Default == nil {
			return p.name + ":long"
		}
		v := p.Default.(int64)
		return fmt.Sprintf("%s:long = long(%d)", p.name, v)
	case types.Real:
		if p.Default == nil {
			return p.name + ":real"
		}
		v := p.Default.(float64)
		return fmt.Sprintf("%s:real = real(%f)", p.name, v)
	case types.String:
		if p.Default == nil {
			return p.name + ":string"
		}
		v := p.Default.(string)
		return fmt.Sprintf(`%s:string = %s`, p.name, utils.QuoteString(v, false))
	case types.Timespan:
		if p.Default == nil {
			return p.name + ":timespan"
		}
		v := p.Default.(time.Duration)
		return fmt.Sprintf("%s:timespan = timespan(%s)", p.name, value.Timespan{Value: v, Valid: true}.Marshal())
	case types.Decimal:
		if p.Default == nil {
			return p.name + ":decimal"
		}

		var sval string
		switch v := p.Default.(type) {
		case string:
			sval = v
		case *big.Float:
			sval = v.String()
		}
		return fmt.Sprintf("%s:decimal = decimal(%s)", p.name, sval)
	}
	panic("internal bug: ParamType.string() called without a call to .validate()")
}
