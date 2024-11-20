package main

import "fmt"

type DeliveryStatus string

const (
	Pending    DeliveryStatus = "Pending"
	Dispatched DeliveryStatus = "Dispatched"
	Delivered  DeliveryStatus = "Delivered"
)

type NewDelivery struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Address string         `json:"address"`
	Status  DeliveryStatus `json:"status"`
}

func createNewDelivery(name string, address string) NewDelivery {
	return NewDelivery{
		ID:      "1",
		Name:    name,
		Address: address,
		Status:  Pending,
	}
}

func main() {
	newDelivery := createNewDelivery("Jena", "Address Test")
	fmt.Println(newDelivery)
}
