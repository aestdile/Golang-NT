
package main

import (
    "fmt"
    "time"
)

func main() {
    go goroutineFunc()
    fmt.Println("Main function is executing...")
    time.Sleep(2 * time.Second)
}


func goroutineFunc() {
    for i := 0; i < 5; i++ {
        fmt.Println("Goroutine is executing...")
        time.Sleep(1 * time.Second)
    }
}


*********************************************************************************************


package main

import ("fmt"
        "time"
)

func main(){
    go display ("Assalomu alaykum...")
    go display ("Vaalaykum assalom...")
    fmt.Scanln()
}

func display(input string){
    for i:=1; true; i++ {
        fmt.Println(i, input)
        time.Sleep(1*time.Second) //break
    }
}


**********************************************************************************************


package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Printf("Worker %d is starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d has finished\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers have finished their jobs")
}


*********************************************************************************************************





























