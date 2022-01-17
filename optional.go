//go:build go1.18

// Copyright (c) 2022 Matrenichev Dmitry
// Under BSD 3-Clause license, see LICENSE file

package optional

import (
	"bytes"
	"encoding/json"
)

// Of creates a new Optional with valid state and provided value.
func Of[T any](val T) Optional[T] {
	return Optional[T]{
		value: val,
		valid: true,
	}
}

// None is helper function which returns empty Optional
func None[T any]() Optional[T] {
	return Optional[T]{}
}

// Optional represents a type T that may be "null" - that, which value is missing.
type Optional[T any] struct {
	value T
	valid bool
}

// Get returns a value, true on non-empty optional and zero-value and false otherwise.
func (o Optional[T]) Get() (T, bool) {
	return o.value, o.valid
}

// ValueOrZero returns a value on non-empty optional and zero-value otherwise.
func (o Optional[T]) ValueOrZero() T {
	return o.value
}

// ValueOr returns a value on non-empty optional or provided value otherwise.
func (o Optional[T]) ValueOr(val T) T {
	if o.valid {
		return o.value
	}

	return val
}

// Valid tells of there any value inside the optional.
func (o Optional[T]) Valid() bool {
	return o.valid
}

// Ptr returns a pointer to value or nil pointer. It is safe to change the value inside the pointer
// (if the T type supports this).
func (o Optional[T]) Ptr() *T {
	if o.valid {
		var store T
		store = o.value
		return &store
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler
func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	*o = None[T]()
	if data == nil || bytes.Equal(data, []byte("null")) {
		*o = Of(*new(T))
		return nil
	}

	if err := json.Unmarshal(data, &o.value); err != nil {
		return err
	}

	o.valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.valid {
		return []byte("null"), nil
	}

	return json.Marshal(o.value)
}
