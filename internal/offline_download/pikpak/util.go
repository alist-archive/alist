package pikpak

import (
	"context"
	"time"

	"codeberg.org/alist/alist/v3/drivers/pikpak"
	"codeberg.org/alist/alist/v3/internal/op"
	"codeberg.org/alist/alist/v3/pkg/singleflight"
	"codeberg.org/alist/go-cache"
)

var taskCache = cache.NewMemCache(cache.WithShards[[]pikpak.OfflineTask](16))
var taskG singleflight.Group[[]pikpak.OfflineTask]

func (p *PikPak) GetTasks(pikpakDriver *pikpak.PikPak) ([]pikpak.OfflineTask, error) {
	key := op.Key(pikpakDriver, "/drive/v1/task")
	if !p.refreshTaskCache {
		if tasks, ok := taskCache.Get(key); ok {
			return tasks, nil
		}
	}
	p.refreshTaskCache = false
	tasks, err, _ := taskG.Do(key, func() ([]pikpak.OfflineTask, error) {
		ctx := context.Background()
		phase := []string{"PHASE_TYPE_RUNNING", "PHASE_TYPE_ERROR", "PHASE_TYPE_PENDING", "PHASE_TYPE_COMPLETE"}
		tasks, err := pikpakDriver.OfflineList(ctx, "", phase)
		if err != nil {
			return nil, err
		}
		// 添加缓存 10s
		if len(tasks) > 0 {
			taskCache.Set(key, tasks, cache.WithEx[[]pikpak.OfflineTask](time.Second*10))
		} else {
			taskCache.Del(key)
		}
		return tasks, nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
