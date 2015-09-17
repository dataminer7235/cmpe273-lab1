package main

import (
  "fmt"
  "time")


func main() {
  //c1 := make(chan string)
  //c2 := make(chan string)

  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f",&input)
  var answer int = fib(int(input))
  val := Time_delay()
  fmt.Println("The value of delay is ")
  fmt.Println(val)
  // select {//case msg1 := <- c1:fmt.Println("Message 1", msg1)
  //         //case msg2 := <- c2:fmt.Println("Message 2", msg2)
  //         case <- time.After(time.Second*5):fmt.Println("The sleep function was in play for the last 5 sec")
  //       }
  fmt.Println("Fibonacci answer after delay is ")
  fmt.Println(answer)
}

func Time_delay()int {
  val_of_sleep := 5
  select {
          case <- time.After(time.Second*5):fmt.Println("The sleep function was in play for the last 500 mili sec")
          return val_of_sleep
        }

}

func fib(x int) int {
  if x == 0 {return 0}
  if x == 1 {return 1}
  return fib(x-1) + fib(x-2)
}
