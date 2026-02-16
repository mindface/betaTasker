package tools

import (
	"strconv"
	dtoquery "github.com/godotask/dto/query"
	"github.com/godotask/interface/http/authcontext" 

	"github.com/gin-gonic/gin"
)

func ParsePagerQuery(c *gin.Context) dtoquery.PagerQuery {
	const (
		defaultPage  = 1
		defaultLimit = 20
		maxPerPage   = 100
	)

	var (
		taskID int
	)
	userID, _ := authcontext.UserID(c)

	page := getPositiveInt(c, "page", defaultPage)
	limit := getPositiveInt(c, "limit", defaultLimit)

	if limit > maxPerPage {
		limit = maxPerPage
	}

	if t := c.Query("task_id"); t != "" {
		if v, err := strconv.Atoi(t); err == nil && v > 0 {
			taskID = v
		}
	}

	return dtoquery.PagerQuery{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
		UserID: userID,
		TaskID: taskID,
	}
}

func getPositiveInt(c *gin.Context, key string, defaultVal int) int {
	if v := c.Query(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			return i
		}
	}
	return defaultVal
}

func getOptionalPositiveInt(c *gin.Context, key string) int {
	if v := c.Query(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			return i
		}
	}
	return 0
}