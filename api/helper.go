package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	MinPage = 1
	MinSize = 10
	MaxSize = 100
)

func getPageAndSize(c *gin.Context) (page, size int) {
	p, ok := c.GetQuery("page")
	if ok {
		page, _ = strconv.Atoi(p)
	}
	s, ok := c.GetQuery("size")
	if ok {
		size, _ = strconv.Atoi(s)
	}
	if page < MinPage {
		page = MinPage
	}
	if size < MinSize {
		size = MinSize
	} else if size > MaxSize {
		size = MaxSize
	}
	return
}
