package main

import "fmt"

func main() {
	statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	stations := map[string][]string{
		"kone":   {"id", "nv", "ut"},
		"ktwo":   {"wa", "id", "mt"},
		"kthree": {"or", "nv", "ca"},
		"kfour":  {"nv", "ut"},
		"kfive":  {"ca", "az"},
	}

	bestStation := ""

	for range statesNeeded {
		for state, station := range stations {
			bestStation = station
		}
	}
	fmt.Println(stations)
}
