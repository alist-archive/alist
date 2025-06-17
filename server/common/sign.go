package common

import (
	stdpath "path"

	"codeberg.org/alist/alist/v3/internal/conf"
	"codeberg.org/alist/alist/v3/internal/model"
	"codeberg.org/alist/alist/v3/internal/setting"
	"codeberg.org/alist/alist/v3/internal/sign"
)

func Sign(obj model.Obj, parent string, encrypt bool) string {
	if obj.IsDir() || (!encrypt && !setting.GetBool(conf.SignAll)) {
		return ""
	}
	return sign.Sign(stdpath.Join(parent, obj.GetName()))
}
