package bootstrap

import (
	"codeberg.org/alist/alist/v3/internal/offline_download/tool"
	"codeberg.org/alist/alist/v3/pkg/utils"
)

func InitOfflineDownloadTools() {
	for k, v := range tool.Tools {
		res, err := v.Init()
		if err != nil {
			utils.Log.Warnf("init tool %s failed: %s", k, err)
		} else {
			utils.Log.Infof("init tool %s success: %s", k, res)
		}
	}
}
