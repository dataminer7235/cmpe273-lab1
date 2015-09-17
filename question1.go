package main

import "fmt"

func main() {
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f",&input)
  var answer int = Fib(int(input))
  fmt.Println(answer)
}

func Fib(x int) int {
  if x == 0 {return 0}
  if x == 1 {return 1}
  return Fib(x-1) + Fib(x-2)
}
