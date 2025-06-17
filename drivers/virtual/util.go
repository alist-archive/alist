package virtual

import (
	"time"

	"codeberg.org/alist/alist/v3/internal/model"
	"codeberg.org/alist/alist/v3/pkg/utils/random"
)

func (d *Virtual) genObj(dir bool) model.Obj {
	obj := &model.Object{
		Name:     random.String(10),
		Size:     0,
		IsFolder: true,
		Modified: time.Now(),
	}
	if !dir {
		obj.Size = random.RangeInt64(d.MinFileSize, d.MaxFileSize)
		obj.IsFolder = false
	}
	return obj
}
