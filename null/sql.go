//go:build go1.18

// Copyright (c) 2022 Matrenichev Dmitry
// Under BSD 3-Clause license, see LICENSE file

package null

import (
	"database/sql"
	"time"

	"github.com/DmitriyMV/optional"
)

// FromNullString is a helper function which converts from sql.NullString to Optional[string]
func FromNullString(val sql.NullString) optional.Optional[string] {
	return valueOrZero(val.Valid, optional.Of(val.String))
}

// FromNullTime is a helper function which converts from sql.NullTime to Optional[time.Time]
func FromNullTime(val sql.NullTime) optional.Optional[time.Time] {
	return valueOrZero(val.Valid, optional.Of(val.Time))
}

// FromNullBool is a helper function which converts from sql.NullBool to Optional[bool]
func FromNullBool(val sql.NullBool) optional.Optional[bool] {
	return valueOrZero(val.Valid, optional.Of(val.Bool))
}

// FromNullByte is a helper function which converts from sql.NullByte to Optional[byte]
func FromNullByte(val sql.NullByte) optional.Optional[byte] {
	return valueOrZero(val.Valid, optional.Of(val.Byte))
}

// FromNullInt16 is a helper function which converts from sql.NullInt16 to Optional[int16]
func FromNullInt16(val sql.NullInt16) optional.Optional[int16] {
	return valueOrZero(val.Valid, optional.Of(val.Int16))
}

// FromNullInt32 is a helper function which converts from sql.NullInt32 to Optional[int32]
func FromNullInt32(val sql.NullInt32) optional.Optional[int32] {
	return valueOrZero(val.Valid, optional.Of(val.Int32))
}

// FromNullInt64 is a helper function which converts from sql.NullInt64 to Optional[int64]
func FromNullInt64(val sql.NullInt64) optional.Optional[int64] {
	return valueOrZero(val.Valid, optional.Of(val.Int64))
}

// FromNullFloat64 is a helper function which converts from sql.NullFloat64 to Optional[float64]
func FromNullFloat64(val sql.NullFloat64) optional.Optional[float64] {
	return valueOrZero(val.Valid, optional.Of(val.Float64))
}

func valueOrZero[T any](ok bool, value T) T {
	if ok {
		return value
	}

	return *new(T)
}
