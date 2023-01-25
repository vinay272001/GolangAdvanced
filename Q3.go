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
