package middlewares

import (
	"strings"

	"codeberg.org/alist/alist/v3/internal/conf"
	"codeberg.org/alist/alist/v3/pkg/utils"
	"codeberg.org/alist/alist/v3/server/common"
	"github.com/gin-gonic/gin"
)

func StoragesLoaded(c *gin.Context) {
	if conf.StoragesLoaded {
		c.Next()
	} else {
		if utils.SliceContains([]string{"", "/", "/favicon.ico"}, c.Request.URL.Path) {
			c.Next()
			return
		}
		paths := []string{"/assets", "/images", "/streamer", "/static"}
		for _, path := range paths {
			if strings.HasPrefix(c.Request.URL.Path, path) {
				c.Next()
				return
			}
		}
		common.ErrorStrResp(c, "Loading storage, please wait", 500)
		c.Abort()
	}
}
