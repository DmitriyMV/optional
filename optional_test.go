//go:build go1.18

// Copyright (c) 2022 Matrenichev Dmitry
// Under BSD 3-Clause license, see LICENSE file

package optional

import (
	"bytes"
	"fmt"
	"testing"
)

func TestOptionalIntGet(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		value Optional[int]
		want  int
		ok    bool
	}{
		"test empty optional": {
			want: 0,
			ok:   false,
		},
		"test none optional": {
			value: None[int](),
			want:  0,
			ok:    false,
		},
		"test valid optional": {
			value: Of(1),
			want:  1,
			ok:    true,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(
			name, func(t *testing.T) {
				t.Parallel()
				val, ok := tt.value.Get()
				if val != tt.want {
					t.Errorf("Get() value got = %v, want %v", val, tt.want)
				}

				if ok != tt.ok {
					t.Errorf("Get() ok got = %v, want %v", val, tt.want)
				}
			},
		)
	}
}

func TestOptionalIntValueOrZero(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		value Optional[int]
		want  int
	}{
		"test empty optional": {
			want: 0,
		},
		"test none optional": {
			value: None[int](),
			want:  0,
		},
		"test valid optional": {
			value: Of(1),
			want:  1,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(
			name, func(t *testing.T) {
				t.Parallel()
				val := tt.value.ValueOrZero()
				if val != tt.want {
					t.Errorf("ValueOrZero() value got = %v, want %v", val, tt.want)
				}
			},
		)
	}
}

func TestOptionalIntValueOr(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		value Optional[int]
		want  int
		or    int
	}{
		"test empty optional": {
			want: 1,
			or:   1,
		},
		"test none optional": {
			value: None[int](),
			want:  1,
			or:    1,
		},
		"test valid optional": {
			value: Of(2),
			want:  2,
			or:    1,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(
			name, func(t *testing.T) {
				t.Parallel()
				val := tt.value.ValueOr(tt.or)
				if val != tt.want {
					t.Errorf("ValueOr() value got = %v, want %v", val, tt.want)
				}
			},
		)
	}
}

func TestOptionalIntValid(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		value Optional[int]
		want  bool
	}{
		"test empty optional": {
			want: false,
		},
		"test none optional": {
			value: None[int](),
			want:  false,
		},
		"test valid optional": {
			value: Of(1),
			want:  true,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(
			name, func(t *testing.T) {
				t.Parallel()
				val := tt.value.Valid()
				if val != tt.want {
					t.Errorf("ValueOr() value got = %v, want %v", val, tt.want)
				}
			},
		)
	}
}

func TestOptionalIntUnmarshalJSON(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		want    Optional[int]
		wantErr bool
		have    []byte
	}{
		"test optional on nil input": {
			want: Of(0),
			have: nil,
		},
		"test optional on `null` input": {
			want: Of(0),
			have: []byte("null"),
		},
		"test optional on `1` input": {
			want: Of(1),
			have: []byte("1"),
		},
		"test optional on `1.2` input": {
			want:    None[int](),
			wantErr: true,
			have:    []byte("1.2"),
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(
			name, func(t *testing.T) {
				t.Parallel()
				var got Optional[int]
				if err := got.UnmarshalJSON(tt.have); (err != nil) != tt.wantErr {
					t.Errorf("UnmarshalJSON() error = %v", err)
				}

				if got != tt.want {
					t.Errorf("Incorrect result of UnmarshalJSON got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestOptionalIntMarshalJSON(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		have    Optional[int]
		want    []byte
		wantErr bool
	}{
		"test output on empty optional": {
			want: []byte("null"),
		},
		"test output on none optional": {
			have: None[int](),
			want: []byte("null"),
		},
		"test output on valid optional": {
			have: Of(2),
			want: []byte("2"),
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(
			name, func(t *testing.T) {
				t.Parallel()
				if got, err := tt.have.MarshalJSON(); (err != nil) != tt.wantErr {
					t.Errorf("MarshalJSON() error = %v", err)
				} else if bytes.Compare(got, tt.want) != 0 {
					t.Errorf("Incorrect result of MarshalJSON got = %v, want %v", string(got), string(tt.want))
				}
			},
		)
	}
}

func TestOptionalIntPtr(t *testing.T) {
	val := Of(2)

	firstPtr := val.Ptr()
	secondPtr := val.Ptr()
	if firstPtr == secondPtr {
		t.Errorf("Returned pointers should not be equal")
	}

	if *firstPtr != *secondPtr {
		t.Errorf("Returned value should be equal")
	}
}

func Example() {
	val := None[int]()
	val = Of(2)
	slc, err := val.MarshalJSON()
	fmt.Println(string(slc), err)

	ptr := val.Ptr()
	fmt.Println(*ptr)
	// Output:
	// 2 <nil>
	// 2
}
