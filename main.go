package main

import (
	"fmt"

	db "github.com/jenagansivakumar/deliverytracking/database"
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

	conn, err := db.OpenDatabase()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}
	defer conn.Close()
	fmt.Println("Database connection successful!")

}
