package middlewares

import (
	"net/url"
	stdpath "path"

	"codeberg.org/alist/alist/v3/internal/errs"
	"codeberg.org/alist/alist/v3/internal/model"
	"codeberg.org/alist/alist/v3/internal/op"
	"codeberg.org/alist/alist/v3/server/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func FsUp(c *gin.Context) {
	path := c.GetHeader("File-Path")
	password := c.GetHeader("Password")
	path, err := url.PathUnescape(path)
	if err != nil {
		common.ErrorResp(c, err, 400)
		c.Abort()
		return
	}
	user := c.MustGet("user").(*model.User)
	path, err = user.JoinPath(path)
	if err != nil {
		common.ErrorResp(c, err, 403)
		return
	}
	meta, err := op.GetNearestMeta(stdpath.Dir(path))
	if err != nil {
		if !errors.Is(errors.Cause(err), errs.MetaNotFound) {
			common.ErrorResp(c, err, 500, true)
			c.Abort()
			return
		}
	}
	if !(common.CanAccess(user, meta, path, password) && (user.CanWrite() || common.CanWrite(meta, stdpath.Dir(path)))) {
		common.ErrorResp(c, errs.PermissionDenied, 403)
		c.Abort()
		return
	}
	c.Next()
}
