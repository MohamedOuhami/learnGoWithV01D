package main

func learning_loops() {
  
  // For loops
  i := 1
  for i <= 3 {
    println(i)
    i++
  }

  // Or we can use the range to set the number of repetitions
  for i := range 3 {
    println(i)
  }
}
