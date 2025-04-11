
package main

import "fmt"

func main() {
    x, y, z := 5, 10, 15
    fmt.Println("Dastlabki values: ", x, y, z)
    swapping(&x, &y, &z)
    fmt.Println("almashtirilgandan keyingi values: ", x, y, z)
}

func swapping(a, b, c *int) {
    *a, *b, *c = *c, *a, *b
}


------------------------------------------------------------------------------


package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "kjfu49ur8hnf3ou4r89hf4no294hrfio4fj"
	natija := nums(str)
	fmt.Println("Numbers: ", natija)
}


func nums(str string) string {
	var a string
	for _, char := range str {
		if unicode.IsDigit(char) {
			a += string(char)
		}
	}
	return a
}


------------------------------------------------------------------------------------


package main

import "fmt"

func main() {
    nums := []int{1, 2, 3}
    slic(nums, 4, 5, 6)
}


func slic(slice []int, values ...int) {
    slice = append(slice, values...)
    slice = append(slice, values...)
    fmt.Println(slice)
}


----------------------------------------------------------------------------------


package main

import "fmt"

func main() {
    nums := []int{5, 3, 7, 1, 9, 2}
    sort_slice(nums)
}

func sort_slice(slice []int) {
    for i := 0; i < len(slice); i++ {
        for j := i + 1; j < len(slice); j++ {
            if slice[i] > slice[j] {
                slice[i], slice[j] = slice[j], slice[i]
            }
        }
    }
    fmt.Println(slice)
}


-----------------------------------------------------------------------------------
































































