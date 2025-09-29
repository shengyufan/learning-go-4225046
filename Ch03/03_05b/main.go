package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")

	// map[typeA]typeB defines a map using typeA param as the key and typeB param as the value
	states := make(map[string]string)
	states["NSW"] = "New South Wales"
	states["QLD"] = "Queensland"
	states["VIC"] = "Victoria"
	fmt.Println(states)

	queensland := states["QLD"]
	fmt.Println(queensland)

	// delete the item
	delete(states, "QLD")
	fmt.Println(states)

	// add the item
	states["WA"] = "Western Australia"
	fmt.Println(states)

	// iterate the map (different order than adding)
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// iterate the map in specified order
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Println("keys sorted")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
