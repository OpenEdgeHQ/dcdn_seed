package dal

import (
	"context"
	"seed.manager/biz/model/dcdn_seed"
	"sync"
)

type task struct {
	sync.RWMutex
	fileInfo dcdn_seed.FileInfo
	num      int32
}

var tasks sync.Map

func AddTask(ctx context.Context, appid int32, fid, url string, max int32) error {
	k := toIndexKey(appid, fid)
	if v, ok := tasks.Load(k); ok {
		t := v.(*task)
		t.Lock()
		defer t.Unlock()

		t.num = max
		t.fileInfo.URL = &url
	} else {
		tasks.Store(k, &task{fileInfo: dcdn_seed.FileInfo{Fid: fid, AppID: appid, URL: &url}, num: max})
	}
	return nil
}

func GetTask(ctx context.Context, peer string, max int32) ([]*dcdn_seed.FileInfo, error) {
	file := make([]*dcdn_seed.FileInfo, 0, max)
	tasks.Range(func(k, v interface{}) bool {
		tm := v.(*task)
		tm.RLock()
		f := tm.fileInfo
		tm.RUnlock()

		if !PeerHasFid(ctx, f.AppID, f.Fid, peer) {
			tm.Lock()
			defer tm.Unlock()

			if tm.num > 0 {
				tm.num--
				file = append(file, &f)
				return len(file) < int(max)
			}
		}
		return true
	})
	return file, nil
}
