// Interfaces 
//       Create an interface to calculate the area of Square and Rectangle
//       type Shape interface{
//              Area() float64
//       }
      
//       type Square struct{
//            Side float64
//       }

//      type Rectangle struct{
//          Height float64
//          Width float64
//      }

//       Implement the Shape interface using type Square and rectangle

package main

import "fmt"

type Shape interface {
    Area() float64
}

type Square struct {
    Side float64
}

type Rectangle struct {
    Height float64
    Width  float64
}

func (s Square) Area() float64 {
    return s.Side * s.Side
}

func (r Rectangle) Area() float64 {
    return r.Height * r.Width
}

func main() {
    s := Square{Side: 5}
    r := Rectangle{Height: 10, Width: 15}

    fmt.Println("Area of Square: ", s.Area())
    fmt.Println("Area of Rectangle: ", r.Area())
}
