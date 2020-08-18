package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ProductID   int    `json:"productId"`
	ProductName string `json:"productName"`
}

func main() {
	product := &Product{
		ProductID:   123,
		ProductName: "pencil",
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(productJSON))
}
