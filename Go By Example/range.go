package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 4 {
			fmt.Println("index:", i)
		}
	}

	mapsgo := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range mapsgo {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range mapsgo {
		fmt.Println("keys:", k)
	}

	// unicode
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
