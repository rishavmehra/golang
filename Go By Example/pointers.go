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

// They are the pointer denotations. The ‘*’ is a pointer to a variable, and ‘&’ is its memory address.

// int a; /* Declares a variable of type integer */
// int *b; /* Declares a pointer to type integer */
// int &c; /* Not possible */

// a = 10; /* a gets the value of 10 */
// b = 10; /* not possible */
// b = &a; /* b gets the address of a */
// *b = 20; /* a now has the value 20 */

// Since the asterisk resembles a gold coin, I read *b as “value of b”.
// Ampersand is also pronounced as ‘et’ which is close enough to “at”, so I read &a as “at a”

// You can set a value for the pointer, but only once you've asked for memory for it using "new". This is how your code should look

// b = new int; /* you create a new integer variable where *b points */
// *b = 20; /* b now has the value 20 */
