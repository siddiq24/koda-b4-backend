package libs

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BuildPageURL(c *gin.Context, page int) string {
	q := c.Request.URL.Query()
	q.Set("page", fmt.Sprintf("%d", page))

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	xfProto := c.GetHeader("X-Forwarded-Proto")
	if xfProto != "" {
		scheme = xfProto
	}

	host := c.GetHeader("X-Forwarded-Host")
	if host == "" {
		host = c.Request.Host
	}
	fmt.Println("URL :", c.FullPath())

	return fmt.Sprintf("%s://%s%s?%s",
		scheme,
		host,
		c.FullPath(),
		q.Encode(),
	)
}
