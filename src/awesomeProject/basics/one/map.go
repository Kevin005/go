package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "dongjiang"
	countryCapitalMap["India"] = "mardrry"
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	/*查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["India"] /*如果确定是真实的,则存在,否则不存在 */
	/*删除元素*/ delete(countryCapitalMap, "France")
	delete(countryCapitalMap, "India")

	if ok {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国首都不存在")
	}

}
