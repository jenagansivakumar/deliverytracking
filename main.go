package main

import "fmt"

type deliveries struct {
	ID           string `json:"id"`
	CustomerName string `json:"customername"`
	Address      string `json:"address"`
}

func main() {
	fmt.Println("Hello world")
}
