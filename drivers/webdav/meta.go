package webdav

import (
	"codeberg.org/alist/alist/v3/internal/driver"
	"codeberg.org/alist/alist/v3/internal/op"
)

type Addition struct {
	Vendor   string `json:"vendor" type:"select" options:"sharepoint,other" default:"other"`
	Address  string `json:"address" required:"true"`
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
	driver.RootPath
	TlsInsecureSkipVerify bool `json:"tls_insecure_skip_verify" default:"false"`
}

var config = driver.Config{
	Name:        "WebDav",
	LocalSort:   true,
	OnlyProxy:   true,
	DefaultRoot: "/",
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &WebDav{}
	})
}
