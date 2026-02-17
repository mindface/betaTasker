package query

// FilterTarget は include に指定できる対象
type FilterTarget string

const (
	FilterUser   FilterTarget = "user"
	FilterTask   FilterTarget = "task"
	FilterMemory FilterTarget = "memory"
	FilterPattern   FilterTarget = "pattern"
	FilterInsight   FilterTarget = "insight"
	FilterModeler FilterTarget = "modeler"
)

// QueryFilter は一覧・検索 API 用の共通 DTO
type QueryFilter struct {
	UserID   *uint
	TaskID   *int
	MemoryID *int

	// include=user,task,memory
	Include []FilterTarget
}

