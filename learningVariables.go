package main

func learnVariables() {

	// This is the print statement in go
	println("Hello World")

	// There are 2 ways to setup the variables
	var name = "Raja"
	println(name)

	// We can also explicitely type the variable
	var b, c int = 5, 6
	println(b, c)

	// When variables are declared and not initialized, it is called zero valued so the ints will take 0 and the strings will be empty
	var d int
	println(d)

	// When declare and initialize the variables using the :=
	fav_team := "Raja Club Athletic"
	println(fav_team)

}
