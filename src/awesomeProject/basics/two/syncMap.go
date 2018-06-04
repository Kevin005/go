package main

import (
	"sync"
	"time"
	"fmt"
)

type SafeMap struct {
	sync.RWMutex
	Map map[int]int
}

func main() {
	safeMap := newSafeMap(10)
	for i := 0;i < 10000 ;i++  {
		go safeMap.readMap(i)
		go safeMap.writeMap(i,i)
	}
	time.Sleep(10 * time.Second)
	a := safeMap.Map
	for k,v := range a{
		fmt.Println("%d===%d",k ,v)
	}
}

func newSafeMap(size int) *SafeMap{
	sm := new(SafeMap)
	sm.Map = make(map[int]int)
	return sm
}

func (sm *SafeMap) readMap(key int) int{
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()
	return value
}

func (sm *SafeMap) writeMap(key int,value int){
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}