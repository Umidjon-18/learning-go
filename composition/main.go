package main

import "fmt"

type Address struct {
	State   string
	City    string
	Street  string
	ZipCode int
}

func (a Address) FullAddress() string {
	if a.State == "" || a.City == "" {
		return "No full address specified"
	}
	return fmt.Sprintf("%s %s %s %d", a.State, a.City, a.Street, a.ZipCode)
}

type Customer struct {
	FirstName       string
	LastName        string
	MainAddress     Address
	ShippingAddress Address
}

func (c Customer) GetCustomerDetails() string {
	return fmt.Sprintf("%s %s %s %s",
		c.FirstName,
		c.LastName,
		c.MainAddress.FullAddress(),
		c.ShippingAddress.FullAddress())
}

func main() {
	customer1 := Customer{
		FirstName: "John",
		LastName:  "Obama",
		MainAddress: Address{
			State:   "USA",
			City:    "Newyork",
			Street:  "Manhettan",
			ZipCode: 121211,
		},
		ShippingAddress: Address{
			State:   "Iran",
			City:    "Tehran",
			Street:  "Khamanaiy",
			ZipCode: 7777777,
		},
	}
	fmt.Println(customer1.MainAddress.FullAddress())
	fmt.Println(customer1.GetCustomerDetails())
}
