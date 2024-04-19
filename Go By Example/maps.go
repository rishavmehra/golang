package main

import (
	"fmt"
	"maps"
)

func main() {

	//To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("maps:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map after delete: ", m)

	clear(m)
	fmt.Println("clear map:", m)

	_, prs := m["k2"]
	fmt.Println("prs", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
