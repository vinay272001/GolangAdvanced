//  Fix deadlock error in the below program :

// func main() {
//    ch := make(chan int)

//    go func() {
//        for i := 0; i < 500; i++ {
//            ch <- i
//        }
//    }()

//    for i := range ch {
//        fmt.Println(i)
//    }

// }
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 500; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
