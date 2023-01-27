// Goroutines
//         Create a struct of employee with data name and age
//                     Create a goroutine which Send the Employees data to the channel with                  incrementing  age
//                     In main function create a buffered channel and then call the 10 goroutines
//                     (use WaitGroups)
                    
// Output :  
//                            {Rizwana, 24}
//                            {Rizwana, 25}
//                            {Rizwana, 26}
//                            {Rizwana, 27}
//                            ……..
//                            {Rizwana, 33}


package main

import (
	"fmt"
	"sync"
)

type employee struct {
	name string
	age  int
}

func response(s employee, wg *sync.WaitGroup) {
	fmt.Println(s.name, " ", s.age)
	wg.Done()
}

func process(s employee, ch chan employee, wg *sync.WaitGroup) {
	ch <- s
	wg.Done()
}
func main() {
	name := "Employee 1"
	var wg sync.WaitGroup

	ch := make(chan employee, 5)
	
	for i := 0; i < 10; i = i + 1 {
		s := employee{
			name: name,
			age:  i,
		}
		wg.Add(1)
		go process(s, ch, &wg)

		wg.Add(1)
		go response(<-ch, &wg)
	}

	wg.Wait()
}
