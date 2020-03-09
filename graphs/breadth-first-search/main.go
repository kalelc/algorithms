package main

import "fmt"

func main() {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
	fmt.Printf("%+v\n", graph)

	var queue []string
	queue = append(queue, graph["you"]...)

	for i, value := range queue {

		fmt.Println(value)
		if i < len(queue) {
			queue = removeElement(queue, i)
		} else {
			queue = nil
		}

	}
}

func removeElement(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
