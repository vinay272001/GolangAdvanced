// Type conversion
// Convert the float value 67.88 into int
// Create 3 variables num1 (int), num2(float64), sum(float64), Add num1 and num2

package main

import (
	"fmt"
)

func main() {
	var f float64 = 67.88
	i := int(f)
	fmt.Println("Float value 67.88 after conversion to int: ", i)

	var num1 int = 10
	var num2 float64 = 5.5
	var sum float64
	sum = float64(num1) + num2
	fmt.Println("Sum of num1 and num2: ", sum)
}
