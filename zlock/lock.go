package zlock

import "sync"

type ZLockMap struct {
	SyncMap sync.Map
}

func New(m sync.Map) *ZLockMap {
	return &ZLockMap{m}
}

func (zm *ZLockMap) Lock(key interface{}) {
	val, ok := zm.SyncMap.Load(key)
	if ok {
		l := val.(*sync.Mutex)
		l.Lock()
	} else {
		l := &sync.Mutex{}
		l.Lock()
		zm.SyncMap.Store(key, l)
	}
}

func (zm *ZLockMap) Unlock(key interface{}) {
	val, ok := zm.SyncMap.Load(key)
	if ok {
		l := val.(*sync.Mutex)
		l.Unlock()
	}
}
