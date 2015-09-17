package main

import "fmt"

type Circle struct {x, y, r float64}
type Rectangle struct{l1,b1 float64}

type Shape interface {
  Area() float64
  Perimeter() float64
}


func main() {



  c := Circle{0,0,5}
  fmt.Println(c.area)
  fmt.Println(c.Perimeter)

  r := Rectangle{6,7}
  fmt.Println(r.area)
  fmt.Println(r.Perimeter)
	/*fmt.Print("Enter a number: ")
	var input float64fmt.Scanf("%f", &input)
	output := input * 2fmt.Println(output)
	*/}

func (c *Circle)area() float64 {
  return float64(3.14 * c.r * c.r)
}
//--------------------------------------
func (c *Circle)Perimeter() float64 {
  return float64(3.14 * c.r * 2)
}

func (r *Rectangle)area() float64{
 return float64(r.l1 * r.b1)
}
//---------------------------------------
func (r *Rectangle)Perimeter() float64 {
  return float64(2 * (r.l1 + r.b1))
}
