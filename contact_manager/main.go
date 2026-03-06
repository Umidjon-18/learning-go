package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contacts []Contact
var contactNameByIndex map[string]int
var nextId int

func init() {
	contacts = make([]Contact, 0)
	contactNameByIndex = make(map[string]int)
	nextId = 0
}

func addContact(name, email, phone string) {
	if _, exists := contactNameByIndex[name]; exists {
		fmt.Println("Contact is already exist")
		return
	}
	newContact := Contact{
		ID:    nextId,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	contacts = append(contacts, newContact)
	contactNameByIndex[name] = nextId
	nextId++
}

func findContact(name string) *Contact {
	index, exists := contactNameByIndex[name]

	if !exists {
		return nil
	}
	return &contacts[index]

}

func listContacts() {
	if len(contacts) == 0 {
		fmt.Println("Contact list is empty")
	} else {
		for index, ctc := range contacts {
			fmt.Printf("ID: %+v Contact: %+v \n", index, ctc)
		}
	}
}

func main() {
	addContact("John", "john@gmail.com", "+998999999999")
	addContact("Bob", "bob@gmail.com", "+998933333333")
	addContact("Jany", "jany@gmail.com", "+998911111111")
	addContact("Alex", "alex@gmail.com", "+998944444444")
	addContact("John", "john@gmail.com", "+998999999999")

	listContacts()

	search := "Joh"
	result := findContact(search)
	if result == nil {
		fmt.Printf("No contact found with name: %+v\n", search)
	} else {
		fmt.Println("Contact found")
	}

}
