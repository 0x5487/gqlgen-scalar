package model

import (
	"fmt"
	"io"
	"strings"
)

type Int64 int64

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (i *Int64) UnmarshalGQL(v interface{}) error {
	val, ok := v.(int64)
	if !ok {
		return fmt.Errorf("value must be int64 type")
	}

	*i = Int64(val)

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (i Int64) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, `"%d"`, i)
}

type Banned bool

func (b Banned) MarshalGQL(w io.Writer) {
	if b {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func (b *Banned) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		*b = strings.ToLower(v) == "true"
		return nil
	case bool:
		*b = Banned(v)
		return nil
	default:
		return fmt.Errorf("%T is not a bool", v)
	}
}
