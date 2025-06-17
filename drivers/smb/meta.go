package smb

import (
	"codeberg.org/alist/alist/v3/internal/driver"
	"codeberg.org/alist/alist/v3/internal/op"
)

type Addition struct {
	driver.RootPath
	Address   string `json:"address" required:"true"`
	Username  string `json:"username" required:"true"`
	Password  string `json:"password"`
	ShareName string `json:"share_name" required:"true"`
}

var config = driver.Config{
	Name:        "SMB",
	LocalSort:   true,
	OnlyLocal:   true,
	DefaultRoot: ".",
	NoCache:     true,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &SMB{}
	})
}
