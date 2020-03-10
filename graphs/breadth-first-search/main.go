package main

import "fmt"

func main() {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "claire", "bob"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	queue := graph["you"]

	peopleChecked := make(map[string]bool)

	for len(queue) != 0 {
		fmt.Print("::Status:: -> ")
		fmt.Println(queue)

		value := queue[0]
		queue = append(queue[:0], queue[1:]...)
		fmt.Println("<< " + value)

		if peopleChecked[value] != true {

			if personIsSeller(value) {
				fmt.Println(value + " is seller!!")
				break
			} else {
				queue = append(queue, graph[value]...)
				fmt.Print(">> ")
				fmt.Printf("%v\n", graph[value])
				peopleChecked[value] = true
			}
		}
	}
}

func personIsSeller(name string) bool {
	if string(name[len(name)-1]) == "m" {
		return true
	} else {
		return false
	}
}
