package kodbox

import (
	"codeberg.org/alist/alist/v3/internal/driver"
	"codeberg.org/alist/alist/v3/internal/op"
)

type Addition struct {
	driver.RootPath

	Address  string `json:"address" required:"true"`
	UserName string `json:"username" required:"false"`
	Password string `json:"password" required:"false"`
}

var config = driver.Config{
	Name:        "KodBox",
	DefaultRoot: "",
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &KodBox{}
	})
}
