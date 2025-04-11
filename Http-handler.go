
----------------------------------------------------------------------------------------------------------------------


package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", SimpleServer)
	http.ListenAndServe(":8080", nil)
}


func SimpleServer(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world, I choose to be happy! %s", r.URL.Path[1:])

}


-------------------------------------------------------------------------------------------------------------------------


package main

import (
	"fmt"
	"net/http"
	"time"
)


func main() {
	http.HandleFunc("/time", timeHandler)

	fmt.Println("Server ishga tushirildi...")
	http.ListenAndServe(":8080", nil)  // natijani olish uchun Google ushbu link orqali murojaat qiling! -->  http://localhost:8080/time
}


func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()

	fmt.Fprintf(w, "Ayni paytdagi vaqt: %s", currentTime.Format("2006-01-02 15:04:05"))

}


-----------------------------------------------------------------------------------------------------------------------------


package main

import (
	"fmt"
	"net/http"
	"time"
)

func lastDayHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()

	if r.Method == "POST" {
		currentTime = currentTime.AddDate(0, 0, -1) 
	}

	lastDay := currentTime.Format("2006-01-02 (Monday)")

	fmt.Fprintf(w, "Kechagi sana va hafta kuni: %s", lastDay)
}

func main() {
	http.HandleFunc("/lastday", lastDayHandler)

	fmt.Println("Server ishga tushirildi...")
	http.ListenAndServe(":8080", nil)  //  natijani olish uchun Google ushbu link orqali murojaat qiling! -->    http://localhost:8080/lastday

}

-----------------------------------------------------------------------------------------------------------------------------------


package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func randomHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) 

	fmt.Fprintf(w, "Tasodifiy raqam: %d", randomNumber)
}

func main() {
	http.HandleFunc("/random", randomHandler)

	fmt.Println("Server ishga tushirildi...")
	http.ListenAndServe(":8080", nil)  // natijani olish uchun Google ushbu link orqali murojaat qiling! -->    http://localhost:8080/random


}


-----------------------------------------------------------------------------------------------------------------------------------------


package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func multiplicationTableHandler(w http.ResponseWriter, r *http.Request) {
	
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			r:=i*j
			fmt.Fprintf(w, strconv.Itoa(i) + "*" + strconv.Itoa(j) + "=" + strconv.Itoa(r) + "\n")
		}
		fmt.Fprintf(w, "\n")
	}
	
}

func main() {
	http.HandleFunc("/mul", multiplicationTableHandler)
	fmt.Println("Server ishga tushirildi...")
	http.ListenAndServe(":8080", nil)  // natijani olish uchun Google ushbu link orqali murojaat qiling! -->    http://localhost:8080/mul


}


--------------------------------------------------------------------------------------------------------------------------------------------------

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	LName string `json:"lName"`
}

func meHandler(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("dune.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var people []Person
	err = json.NewDecoder(file).Decode(&people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {

	http.HandleFunc("/me", meHandler)

	fmt.Println("Server ishga tushirildi...")
	http.ListenAndServe(":8080", nil)  // natijani olish uchun Google ushbu link orqali murojaat qiling! -->    http://localhost:8080/me

}

Error:   Bunday fayl mavjud emas deya xato beryapti!? 

------------------------------------------------------------------------------------------------------------------------------

















