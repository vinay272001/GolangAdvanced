// Pointers :
// Declare a variable f of type float and assign the value as 4.5
//             Find the address of f and the value of f using address and print it
//             Then update the value of f using address of f

package main

import "fmt"

func main() {
	var f float64 = 4.5
	fmt.Println("Value of f: ", f)
	fmt.Println("Address of f: ", &f)

	ptr := &f
	fmt.Println("Value of f using address: ", *ptr)

	*ptr = 3.1415
	fmt.Println("Updated value of f: ", f)
}
