package query

import (
	"strings"

	querydto "github.com/godotask/dto/query"
	"fmt"
)

// ParseIncludeParam
// "user,task,memory" → []FilterTarget
func ParseIncludeParam(raw string) []querydto.FilterTarget {
	if raw == "" {
		return nil
	}

	parts := strings.Split(raw, ",")
	fmt.Printf("%#v\n", parts)
	fmt.Printf("type=%T value=%v\n", parts, parts)
	includes := make([]querydto.FilterTarget, 0, len(parts))

	for _, p := range parts {
		v := strings.TrimSpace(p)
		if v != "" {
			includes = append(includes, querydto.FilterTarget(v))
		}
	}
	return includes
}
