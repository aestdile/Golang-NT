



package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  int}

type Task struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	file, err := os.Open("13.json")
	if err != nil {
		log.Fatalf("Faylni ochishda xatolik: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var tasks []Task
	if err := decoder.Decode(&tasks); err != nil {
		log.Fatalf("Task JSON unmarshal xatoligi: %v", err)
	}

	fmt.Println("Vazifalar:")
	for _, task := range tasks {
		fmt.Println("User ID:", task.UserID)
		fmt.Println("ID:", task.ID)
		fmt.Println("Title:", task.Title)
		fmt.Println("Completed:", task.Completed)
		fmt.Println()
	}
}





























