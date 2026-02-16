package query

// FilterTarget は include に指定できる対象
type FilterTarget string

const (
	FilterUser   FilterTarget = "user"
	FilterTask   FilterTarget = "task"
	FilterMemory FilterTarget = "memory"
)

// QueryFilter は一覧・検索 API 用の共通 DTO
type QueryFilter struct {
	UserID   *uint
	TaskID   *int
	MemoryID *int

	// include=user,task,memory
	Include []FilterTarget
}
