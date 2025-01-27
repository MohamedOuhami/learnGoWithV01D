package main

import "fmt"

func learn_collections() {
	// An array is a numbered sequences of elemenst with a predefined length
	var a [5]int

	fmt.Println(a)

	fmt.Println("The array has a length of ", len(a))

	// To declare an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// We can have the compiler count the number of elements using ...
	c := [...]int{1, 2, 3}
	fmt.Println(c)
	fmt.Println(len(c))

	d := [...]int{100, 3: 500, 500}
	// It means the element at index 3 will be set to 400 and the rest is 0
	fmt.Println(d)
}
