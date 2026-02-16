package query

type PagerQuery struct {
	Page   int
	Limit  int
	Offset int

	// optional filters
	UserID uint
	TaskID int
}
