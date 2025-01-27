package main

func learn_conditions() {

	if 7%2 == 0 {
		println("7 is an even number")
	} else {
		println("7 is odd")
	}

	// A statement can precede the conditional
	if num := 9; num < 0 {
		println(num, " is negative")
	} else if num > 0 {
		println(num, " is positive")
	} else {
		println(num, " is nil")
	}

	// We can also use switch statements
	i := 2
	print("Write ", i, " as ")

	switch i {
	case 1:
		println("One")
	case 2:
		println("Two")
	case 3:
		println("Three")
	default:
		println("Not one of those 3")
	}

	// A type switch statement compares types instead of values
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			println("This is a bool value")
		case int:
			println("This is an int")
		default:
			println("This is not a bool or int", t)
		}
	}

	whatAmI(true)
	whatAmI(2)
	whatAmI("Hello")
}
