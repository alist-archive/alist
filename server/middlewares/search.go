package middlewares

import (
	"codeberg.org/alist/alist/v3/internal/conf"
	"codeberg.org/alist/alist/v3/internal/errs"
	"codeberg.org/alist/alist/v3/internal/setting"
	"codeberg.org/alist/alist/v3/server/common"
	"github.com/gin-gonic/gin"
)

func SearchIndex(c *gin.Context) {
	mode := setting.GetStr(conf.SearchIndex)
	if mode == "none" {
		common.ErrorResp(c, errs.SearchNotAvailable, 500)
		c.Abort()
	} else {
		c.Next()
	}
}
