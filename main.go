package main

import (
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
)

type DeliveryStatus string

const (
	Pending    DeliveryStatus = "Pending"
	Dispatched DeliveryStatus = "Dispatched"
	Delivered  DeliveryStatus = "Delivered"
)

type Delivery struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Address string         `json:"address"`
	Status  DeliveryStatus `json:"status"`
}

func createNewDelivery(name string, address string) Delivery {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return Delivery{
		ID:      id.String(),
		Name:    name,
		Address: address,
		Status:  Pending,
	}
}

func main() {
	newDelivery := createNewDelivery("jena", "address")
	fmt.Println(newDelivery)
}
