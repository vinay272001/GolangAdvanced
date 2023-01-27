// Maps
//      -      Create an map for storing 5 student marks (type : map[string]int{}
//      -      Check if any student have 34 marks increment it to 1 and then print the map(Iterate over the maps)

package main

import "fmt"

func main() {
	studentMarks := map[string]int{
		"student1": 45,
		"student2": 34,
		"student3": 67,
		"student4": 88,
		"student5": 99,
	}

	for key, value := range studentMarks {
		if value == 34 {
			studentMarks[key] = value + 1
		}
	}

	fmt.Println("Student marks after increment:")
	for key, value := range studentMarks {
		fmt.Println(key, value)
	}
}
