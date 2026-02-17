package tools

func BuildPageMeta(total int64, page, limit int) map[string]interface{} {
	totalPages := 0
	if total > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	return map[string]interface{}{
		"total":       total,
		"total_pages": totalPages,
		"page":        page,
		"limit":       limit,
	}
}
