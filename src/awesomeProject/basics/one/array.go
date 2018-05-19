package main

func main() {
	countryCapitalMap := make(map[string]string)

	countryCapitalMap[""] = ""
	delete(countryCapitalMap, "France")
	delete(countryCapitalMap, "a")
}
