


package main

import "fmt"

func main(){

	// Login="Alice" && Password="Lena_03"

	i:=0
	for {
		i+=1
		var login string
		fmt.Println("Login: ")
		fmt.Scan(&login)

		var password string
		fmt.Println("Password: ")
		fmt.Scan(&password)

		if login=="Alice" && password=="Lena_03"{
			fmt.Println("Succesfull. Welcome here!")
			break
		}else if i==3{
			fmt.Println("Sizdagi urinishlar soni tugadi. Iltimos, birozdan keyin yana urinib ko'ring!")
			break

		}else{
			fmt.Println("Login va Password xato kiritildi. Iltimos qaytadan kiriting: ")
		}
	}

}

----------------------------------------------------------------------------



package main

import ("fmt"
		"math/rand")
func main(){

	
	slc:=[]int{}
	slc2:=[]int{}

	for i:=0; i<10; i++{
		c:=rand.Intn(100)
		v:=rand.Intn(100)
		slc = append(slc, c)
		slc2 = append(slc2, v)
	}

	fmt.Println(slc)
	fmt.Println(slc2)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if slc[i] == slc2[j] {
				fmt.Println(slc[i])
			}
		}
	}

}


-----------------------------------------------------------------------



package main

import ("fmt"
		"math/rand")
func main(){

	
	slc:=[]int{}
	var k int
	fmt.Println("k: ")	
	fmt.Scan(&k)
	c:=0

	for i:=0; i<10; i++{
		c:=rand.Intn(100)
		slc = append(slc, c)
	}
	
	fmt.Println(slc)

	for i:=0; i<10; i++{
		if slc[i]<k{
			c+=1
		}
	}
	fmt.Println("k ga teng elementlar soni: ", c)

}


-----------------------------------------------------------------------

package main

import "fmt"

func main() {
	s := "Hello world, salom dunyo"
	for i := len(s)-1; i >= 0; i-- {
		fmt.Printf(string(s[i]))
	}
}


-------------------------------------------------------------------------




















































