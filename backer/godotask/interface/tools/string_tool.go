package tools

import (
	"strconv"
)

func NullableIntToString(v string) *int {
	if v == "" {
		return nil
	}
	if i, err := strconv.Atoi(v); err == nil {
		return &i
	}
	return nil
}
