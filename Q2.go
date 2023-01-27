// Arrays and slices :
// Create an array of length 5 with values 45, 46, 67, 88, 99 add the values of on the index 2 and 4 and print the result
// Create an array as above and print the slice a[i+3, j], where i is starting index and j is last index

package main

import "fmt"

func main() {
	arr := [5]int{45, 46, 67, 88, 99}
	result := arr[2] + arr[4]
	fmt.Println("Result: ", result)

	// slice a[i+3, j]
	i := 2
	j := 4
	fmt.Println("Slice a[i+3, j]: ", arr[i+3:j])
}
