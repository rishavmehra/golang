package main

import "fmt"

func zeroval(ival int) {
	ival = 0
	// fmt.Println(&ival)
}

func zeroptr(iptr *int) {
	*iptr = 0
	// fmt.Println(&iptr)
}

func main() {
	i := 1
	fmt.Println("initial:", i, &i)

	zeroval(i)
	fmt.Println("zeroval:", i, &i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i, &i)

	fmt.Println("pointer:", &i)

	var num int = 42
	var ptr *int // declaring a pointer to an integer
	ptr = &num   //assigning the address of 'num' to 'ptr'
	// fmt.Println(ptr)
}
