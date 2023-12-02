package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (c Product) GetFullInfo() string {
	return fmt.Sprintf(
		"Name: %s,\nDescription: %s,\nPrice: %d,\n",
		c.Name, c.Description, c.Price,
	)
}

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
		time.Sleep(time.Second)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	course := Product{
		Name:        "Golang Course",
		Description: "Learn Golang fast and deploy scalable projects",
		Price:       100,
	}

	json.NewEncoder(w).Encode(course)
}

func main() {
	// http.HandleFunc("/", home)
	// http.ListenAndServe(":8080", nil)
	// fmt.Println("Server listening on localhost:8080")

	//  counter() // simple execution => prints 3 times 1 -> 10
	//  counter() // running one after another
	//  counter() // takes 30 seconds

	go counter() // go routine => prints 3 times the same counter value
	go counter() // because it's running simultaneously on threads
	counter()    // takes 10 seconds
}
