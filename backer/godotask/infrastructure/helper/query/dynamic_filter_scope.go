package query

import (
	"strings"

	querydto "github.com/godotask/dto/query"
)

// ParseIncludeParam
// "user,task,memory" → []FilterTarget
func ParseIncludeParam(raw string) []querydto.FilterTarget {
	if raw == "" {
		return nil
	}

	parts := strings.Split(raw, ",")
	result := make([]querydto.FilterTarget, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)
		switch querydto.FilterTarget(p) {
		case querydto.FilterUser,
			querydto.FilterTask,
			querydto.FilterMemory:
			result = append(result, querydto.FilterTarget(p))
		}
	}
	return result
}

