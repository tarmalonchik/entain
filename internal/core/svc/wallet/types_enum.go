// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package wallet

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// TransactionSourceTypeGame is a TransactionSourceType of type game.
	TransactionSourceTypeGame TransactionSourceType = "game"
	// TransactionSourceTypeServer is a TransactionSourceType of type server.
	TransactionSourceTypeServer TransactionSourceType = "server"
	// TransactionSourceTypePayment is a TransactionSourceType of type payment.
	TransactionSourceTypePayment TransactionSourceType = "payment"
)

var ErrInvalidTransactionSourceType = errors.New("not a valid TransactionSourceType")

// String implements the Stringer interface.
func (x TransactionSourceType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x TransactionSourceType) IsValid() bool {
	_, err := ParseTransactionSourceType(string(x))
	return err == nil
}

var _TransactionSourceTypeValue = map[string]TransactionSourceType{
	"game":    TransactionSourceTypeGame,
	"server":  TransactionSourceTypeServer,
	"payment": TransactionSourceTypePayment,
}

// ParseTransactionSourceType attempts to convert a string to a TransactionSourceType.
func ParseTransactionSourceType(name string) (TransactionSourceType, error) {
	if x, ok := _TransactionSourceTypeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _TransactionSourceTypeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return TransactionSourceType(""), fmt.Errorf("%s is %w", name, ErrInvalidTransactionSourceType)
}

var errTransactionSourceTypeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

var sqlIntTransactionSourceTypeMap = map[int64]TransactionSourceType{
	0: TransactionSourceTypeGame,
	1: TransactionSourceTypeServer,
	2: TransactionSourceTypePayment,
}

var sqlIntTransactionSourceTypeValue = map[TransactionSourceType]int64{
	TransactionSourceTypeGame:    0,
	TransactionSourceTypeServer:  1,
	TransactionSourceTypePayment: 2,
}

func lookupSqlIntTransactionSourceType(val int64) (TransactionSourceType, error) {
	x, ok := sqlIntTransactionSourceTypeMap[val]
	if !ok {
		return x, fmt.Errorf("%v is not %w", val, ErrInvalidTransactionSourceType)
	}
	return x, nil
}

// Scan implements the Scanner interface.
func (x *TransactionSourceType) Scan(value interface{}) (err error) {
	if value == nil {
		*x = TransactionSourceType("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x, err = lookupSqlIntTransactionSourceType(v)
	case string:
		*x, err = ParseTransactionSourceType(v)
	case []byte:
		if val, verr := strconv.ParseInt(string(v), 10, 64); verr == nil {
			*x, err = lookupSqlIntTransactionSourceType(val)
		} else {
			// try parsing the value as a string
			*x, err = ParseTransactionSourceType(string(v))
		}
	case TransactionSourceType:
		*x = v
	case int:
		*x, err = lookupSqlIntTransactionSourceType(int64(v))
	case *TransactionSourceType:
		if v == nil {
			return errTransactionSourceTypeNilPtr
		}
		*x = *v
	case uint:
		*x, err = lookupSqlIntTransactionSourceType(int64(v))
	case uint64:
		*x, err = lookupSqlIntTransactionSourceType(int64(v))
	case *int:
		if v == nil {
			return errTransactionSourceTypeNilPtr
		}
		*x, err = lookupSqlIntTransactionSourceType(int64(*v))
	case *int64:
		if v == nil {
			return errTransactionSourceTypeNilPtr
		}
		*x, err = lookupSqlIntTransactionSourceType(int64(*v))
	case float64: // json marshals everything as a float64 if it's a number
		*x, err = lookupSqlIntTransactionSourceType(int64(v))
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errTransactionSourceTypeNilPtr
		}
		*x, err = lookupSqlIntTransactionSourceType(int64(*v))
	case *uint:
		if v == nil {
			return errTransactionSourceTypeNilPtr
		}
		*x, err = lookupSqlIntTransactionSourceType(int64(*v))
	case *uint64:
		if v == nil {
			return errTransactionSourceTypeNilPtr
		}
		*x, err = lookupSqlIntTransactionSourceType(int64(*v))
	case *string:
		if v == nil {
			return errTransactionSourceTypeNilPtr
		}
		*x, err = ParseTransactionSourceType(*v)
	default:
		return errors.New("invalid type for TransactionSourceType")
	}

	return
}

// Value implements the driver Valuer interface.
func (x TransactionSourceType) Value() (driver.Value, error) {
	val, ok := sqlIntTransactionSourceTypeValue[x]
	if !ok {
		return nil, ErrInvalidTransactionSourceType
	}
	return int64(val), nil
}

type NullTransactionSourceType struct {
	TransactionSourceType TransactionSourceType
	Valid                 bool
}

func NewNullTransactionSourceType(val interface{}) (x NullTransactionSourceType) {
	err := x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	_ = err            // make any errcheck linters happy
	return
}

// Scan implements the Scanner interface.
func (x *NullTransactionSourceType) Scan(value interface{}) (err error) {
	if value == nil {
		x.TransactionSourceType, x.Valid = TransactionSourceType(""), false
		return
	}

	err = x.TransactionSourceType.Scan(value)
	x.Valid = (err == nil)
	return
}

// Value implements the driver Valuer interface.
func (x NullTransactionSourceType) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	// driver.Value accepts int64 for int values.
	return string(x.TransactionSourceType), nil
}

const (
	// TransactionStateTypeWin is a TransactionStateType of type win.
	TransactionStateTypeWin TransactionStateType = "win"
	// TransactionStateTypeLose is a TransactionStateType of type lose.
	TransactionStateTypeLose TransactionStateType = "lose"
)

var ErrInvalidTransactionStateType = errors.New("not a valid TransactionStateType")

// String implements the Stringer interface.
func (x TransactionStateType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x TransactionStateType) IsValid() bool {
	_, err := ParseTransactionStateType(string(x))
	return err == nil
}

var _TransactionStateTypeValue = map[string]TransactionStateType{
	"win":  TransactionStateTypeWin,
	"lose": TransactionStateTypeLose,
}

// ParseTransactionStateType attempts to convert a string to a TransactionStateType.
func ParseTransactionStateType(name string) (TransactionStateType, error) {
	if x, ok := _TransactionStateTypeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _TransactionStateTypeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return TransactionStateType(""), fmt.Errorf("%s is %w", name, ErrInvalidTransactionStateType)
}

var errTransactionStateTypeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

var sqlIntTransactionStateTypeMap = map[int64]TransactionStateType{
	0: TransactionStateTypeWin,
	1: TransactionStateTypeLose,
}

var sqlIntTransactionStateTypeValue = map[TransactionStateType]int64{
	TransactionStateTypeWin:  0,
	TransactionStateTypeLose: 1,
}

func lookupSqlIntTransactionStateType(val int64) (TransactionStateType, error) {
	x, ok := sqlIntTransactionStateTypeMap[val]
	if !ok {
		return x, fmt.Errorf("%v is not %w", val, ErrInvalidTransactionStateType)
	}
	return x, nil
}

// Scan implements the Scanner interface.
func (x *TransactionStateType) Scan(value interface{}) (err error) {
	if value == nil {
		*x = TransactionStateType("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x, err = lookupSqlIntTransactionStateType(v)
	case string:
		*x, err = ParseTransactionStateType(v)
	case []byte:
		if val, verr := strconv.ParseInt(string(v), 10, 64); verr == nil {
			*x, err = lookupSqlIntTransactionStateType(val)
		} else {
			// try parsing the value as a string
			*x, err = ParseTransactionStateType(string(v))
		}
	case TransactionStateType:
		*x = v
	case int:
		*x, err = lookupSqlIntTransactionStateType(int64(v))
	case *TransactionStateType:
		if v == nil {
			return errTransactionStateTypeNilPtr
		}
		*x = *v
	case uint:
		*x, err = lookupSqlIntTransactionStateType(int64(v))
	case uint64:
		*x, err = lookupSqlIntTransactionStateType(int64(v))
	case *int:
		if v == nil {
			return errTransactionStateTypeNilPtr
		}
		*x, err = lookupSqlIntTransactionStateType(int64(*v))
	case *int64:
		if v == nil {
			return errTransactionStateTypeNilPtr
		}
		*x, err = lookupSqlIntTransactionStateType(int64(*v))
	case float64: // json marshals everything as a float64 if it's a number
		*x, err = lookupSqlIntTransactionStateType(int64(v))
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errTransactionStateTypeNilPtr
		}
		*x, err = lookupSqlIntTransactionStateType(int64(*v))
	case *uint:
		if v == nil {
			return errTransactionStateTypeNilPtr
		}
		*x, err = lookupSqlIntTransactionStateType(int64(*v))
	case *uint64:
		if v == nil {
			return errTransactionStateTypeNilPtr
		}
		*x, err = lookupSqlIntTransactionStateType(int64(*v))
	case *string:
		if v == nil {
			return errTransactionStateTypeNilPtr
		}
		*x, err = ParseTransactionStateType(*v)
	default:
		return errors.New("invalid type for TransactionStateType")
	}

	return
}

// Value implements the driver Valuer interface.
func (x TransactionStateType) Value() (driver.Value, error) {
	val, ok := sqlIntTransactionStateTypeValue[x]
	if !ok {
		return nil, ErrInvalidTransactionStateType
	}
	return int64(val), nil
}

type NullTransactionStateType struct {
	TransactionStateType TransactionStateType
	Valid                bool
}

func NewNullTransactionStateType(val interface{}) (x NullTransactionStateType) {
	err := x.Scan(val) // yes, we ignore this error, it will just be an invalid value.
	_ = err            // make any errcheck linters happy
	return
}

// Scan implements the Scanner interface.
func (x *NullTransactionStateType) Scan(value interface{}) (err error) {
	if value == nil {
		x.TransactionStateType, x.Valid = TransactionStateType(""), false
		return
	}

	err = x.TransactionStateType.Scan(value)
	x.Valid = (err == nil)
	return
}

// Value implements the driver Valuer interface.
func (x NullTransactionStateType) Value() (driver.Value, error) {
	if !x.Valid {
		return nil, nil
	}
	// driver.Value accepts int64 for int values.
	return string(x.TransactionStateType), nil
}
