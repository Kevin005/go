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

	captial, ok := countryCapitalMap["India"]
	if ok {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国首都不存在")
	}

}
