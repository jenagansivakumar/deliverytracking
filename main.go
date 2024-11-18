package main

import "fmt"

type DeliveryStatus string

const (
	Pending    DeliveryStatus = "Pending"
	Dispatched DeliveryStatus = "Dispatched"
	Delivered  DeliveryStatus = "Delivered"
)

type deliveries struct {
	ID           string         `json:"id"`
	CustomerName string         `json:"customername"`
	Address      string         `json:"address"`
	Status       DeliveryStatus `json:"status"`
}

func main() {
	fmt.Println("Hello world")
}
