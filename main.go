package main

import (
	"fmt"
	"net/http"
)

type Product struct {
	Name        string
	Descprition string
	Price       int
}

func (c Product) GetFullInfo() string {
	return fmt.Sprintf(
		"Name: %s,\nDescription: %s,\nPrice: %d,\n",
		c.Name, c.Descprition, c.Price,
	)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	product := Product{
		Name:        "Golang Course",
		Descprition: "Learn Golang fast and deploy scalable projects",
		Price:       100,
	}

	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server listening on localhost:8080")
	fmt.Println(product.GetFullInfo())
}
