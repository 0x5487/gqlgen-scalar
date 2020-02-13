package model

import (
	"fmt"
	"io"
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
