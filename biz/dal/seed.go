package dal

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"seed.manager/biz/model/dcdn_seed"
	"strconv"
	"strings"
	"sync"
	"time"
)

var indexMap sync.Map
var peerInfo sync.Map
var peerFileList sync.Map

type PeerInfo struct {
	Address string `json:"address"`
}

func toIndexKey(appid int32, Fid string) string {
	return fmt.Sprintf("%d_%s", appid, Fid)
}

func parseIndexKey(key string) (appid int32, Fid string, err error) {
	data := strings.Split(key, "_")
	if len(data) != 2 {
		return 0, "", errors.New("invalid index key")
	}

	var aid int64
	aid, err = strconv.ParseInt(data[0], 10, 32)
	if err != nil {
		return 0, "", errors.New("invalid index key")
	}
	appid = int32(aid)
	return appid, data[1], nil
}

func AddPeerInfo(ctx context.Context, peer string, info *PeerInfo) error {
	logger.CtxInfof(ctx, "add peer=%s info=%+v", peer, info)
	peerInfo.Store(peer, info)
	return nil
}

func AddSeedInfo(ctx context.Context, peer string, files []*dcdn_seed.FileInfo) error {
	m, ok := peerFileList.Load(peer)
	if !ok {
		m = sync.Map{}
		peerFileList.Store(peer, m)
	}

	peerFidList := m.(sync.Map)
	for _, file := range files {
		logger.CtxInfof(ctx, "peer=%s fid=%s appid=%d add seedindex", peer, file.Fid, file.AppID)
		k := toIndexKey(file.AppID, file.Fid)
		peerFidList.Store(k, time.Now().Unix())
		list, ok := indexMap.Load(k)
		if !ok {
			l := sync.Map{}
			l.Store(peer, time.Now().Unix())
			indexMap.Store(k, l)
		}

		l, ok := list.(sync.Map)
		if !ok {
			panic("error code type ")
		}
		l.Store(peer, time.Now().Unix())
	}
	return nil
}

func GetALlSeedInfo(peer string) ([]*dcdn_seed.FileInfo, error) {
	m, ok := peerFileList.Load(peer)
	if !ok {
		return nil, nil
	}

	peerFidList := m.(sync.Map)
	ret := make([]*dcdn_seed.FileInfo, 0, 10)
	peerFidList.Range(func(k, v interface{}) bool {
		appid, fid, err := parseIndexKey(k.(string))
		if err != nil {
			return false
		}
		ret = append(ret, &dcdn_seed.FileInfo{
			AppID: appid,
			Fid:   fid,
		})
		return true
	})
	return ret, nil
}

func RemoveSeedInfo(ctx context.Context, peer string, appid int32, fid string) error {
	logger.CtxInfof(ctx, "remove peer=%s appid=%d fid=%s", peer, appid, fid)
	m, ok := peerFileList.Load(peer)
	if !ok {
		return nil
	}

	k := toIndexKey(appid, fid)
	peerFidList := m.(sync.Map)
	peerFidList.Delete(k)

	m2, ok := indexMap.Load(k)
	if !ok {
		return nil
	}

	peerList := m2.(sync.Map)
	peerList.Delete(peer)
	return nil
}

func QueryFidPeers(ctx context.Context, appid int32, Fid string) ([]string, error) {
	ret := make([]string, 0, 10)

	key := toIndexKey(appid, Fid)
	list, ok := indexMap.Load(key)
	if !ok {
		return ret, nil
	}

	m, ok := list.(sync.Map)
	if !ok {
		return ret, nil
	}

	m.Range(func(key, value any) bool {
		ret = append(ret, key.(string))
		return true
	})

	return ret, nil
}

func GetPeerInfo(peer string) (*PeerInfo, error) {
	info, ok := peerInfo.Load(peer)
	if !ok {
		return &PeerInfo{}, errors.New("peer not exist")
	}
	return info.(*PeerInfo), nil
}

func PeerHasFid(ctx context.Context, appid int32, Fid string, peer string) bool {
	k := toIndexKey(appid, Fid)
	peerList, ok := indexMap.Load(k)
	if !ok {
		return false
	}

	peerM := peerList.(sync.Map)
	_, ok = peerM.Load(peer)
	return ok
}
