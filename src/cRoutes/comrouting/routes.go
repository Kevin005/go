package comrouting

import (
	"strings"
	"strconv"
	"fmt"
)

type CRoute_S struct {
	routeTable map[int]map[int]bool
}

func CreateCRouter() *CRoute_S {
	var route CRoute_S
	route.routeTable = make(map[int]map[int]bool)
	//init 32 of routeTable length
	for maskLen := 0; maskLen <= 32; maskLen++ {
		route.routeTable[maskLen] = make(map[int]bool)
	}
	return &route
}

func AddCRouteItem(cr *CRoute_S, prefix string, maskLen int) bool {
	if nil == cr || nil == cr.routeTable {
		return false
	}
	rtVal := ipStr2Val(prefix)
	cr.routeTable[maskLen][rtVal] = true
	return true
}

func DelCRouteItem(cr *CRoute_S, prefix string, maskLen int) bool {
	if nil == cr || nil == cr.routeTable {
		return false
	}
	rtVal := ipStr2Val(prefix)
	if 0 > rtVal {
		return false
	}
	if maskLen < 0 {
		maskLen = 0
	} else if maskLen > 32 {
		maskLen = 32
	}
	rtVal = ipValMasked(rtVal, maskLen)
	if _, bExist := cr.routeTable[maskLen][rtVal]; false == bExist {
		return true
	}
	delete(cr.routeTable[maskLen], rtVal)
	return true
}

func SearchCRoute(cr *CRoute_S, ip string) bool {
	if nil == cr || nil == cr.routeTable {
		return false
	}
	rtVal := ipStr2Val(ip)
	if 0 > rtVal {
		return false
	}
	var maskLen int = 0
	for maskLen = 32; maskLen >= 0; maskLen-- {
		tmpVal := ipValMasked(rtVal, maskLen)
		if _, bExist := cr.routeTable[maskLen][tmpVal]; bExist {
			fmt.Printf("SearchCRoute ip=%s match Item %s\n", ip, ipVal2Str(tmpVal))
			return true
		}
	}
	fmt.Printf("SearchCRoute ip=%s match Item fail\n", ip)
	return false
}

func PrintCRoute(cr *CRoute_S) {
	if nil == cr || nil == cr.routeTable {
		return
	}
	var maskLen int = 0
	for maskLen = 32; maskLen >= 0; maskLen-- {
		if 0 == len(cr.routeTable[maskLen]) {
			continue
		}
		fmt.Printf("maskLen=%d List As:\n", maskLen)
		itemSeq := 1
		for ipVal, _ := range cr.routeTable[maskLen] {
			fmt.Printf("%d: %s\n", itemSeq, ipVal2Str(ipVal))
			itemSeq++
		}
	}
}

func ipStr2Val(ip string) int {
	segs := strings.Split(strings.TrimSpace(ip), ".")
	if 4 != len(segs) {
		return -1
	}
	ret := make([]int, 0)
	for _, segx := range segs {
		val, err := strconv.Atoi(segx) //string 转为 int
		if nil != err {
			return -1
		}
		ret = append(ret, val)
	}
	rtVal := ret[0]<<24 + ret[1]<<16 + ret[2]<<8 + ret[3] //todo ?
	return rtVal
}

func ipVal2Str(ipVal int) string {
	s1 := (ipVal & 0xff000000) >> 24
	s2 := (ipVal & 0xff0000) >> 16
	s3 := (ipVal & 0xff00) >> 8
	s4 := (ipVal & 0xff)
	return fmt.Sprintf("%d.%d.%d.%d", s1, s2, s3, s4)
}

func ipValMasked(ipVal int, maskLen int) int {
	rtVal := (ipVal >> uint(32-maskLen)) & 0xffffffff
	rtVal = rtVal << uint(32-maskLen)
	return rtVal
}
