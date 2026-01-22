package context

import "github.com/gin-gonic/gin"

func GetUserIDFromContext(c *gin.Context) (uint, bool) {
	v, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}

	userID, ok := v.(uint)
	if !ok {
		return 0, false
	}

	return userID, true
}
