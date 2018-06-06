package sync

import "sync"

type SafeMap struct {
	sync.RWMutex
	Map map[int]int
}

func main() {
	safeMap := newSafeMap()

	for i := 0; i < 100000; i++ {
		go safeMap.writeMap(i, i)
		go safeMap.readMap(i)
	}

}

func newSafeMap() *SafeMap {
	sm := new(SafeMap)
	sm.Map = make(map[int]int)
	return sm

}

func (sm *SafeMap) readMap(key int) int {
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()
	return value
}

func (sm *SafeMap) writeMap(key int, value int) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}