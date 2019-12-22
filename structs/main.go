package main

import "fmt"

//Struct properties default values

// string => ""
// int    => 0
// float  => 0
// bool   => false
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	fmt.Println("Working with structs!")
	tasos := person{
		firstName: "Anastasios",
		lastName:  "Theodoisou",
		contactInfo: contactInfo{
			email:   "anastasios.theodosiou@gmail.com",
			zipCode: 54621,
		},
	}

	//Pointer
	tasosPointer := &tasos
	tasosPointer.updateName("Tasos")
	tasos.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func main() {
// 	fmt.Println("Working with structs!")
// 	//First type of declaration
// 	//This approach depends on the order we list out our properties...

// 	// tasos := person{"Anastasios", "Theodosiou"}
// 	// fmt.Println("Hello ", tasos.firstName+" "+tasos.lastName)

// 	//Second type of declaration

// 	// tasos := person{firstName: "Anastasios", lastName: "Theodosiou"}
// 	// fmt.Println("Hello ", tasos.firstName+" "+tasos.lastName)

// 	//Third way
// 	// var tasos person
// 	// tasos.firstName = "Anastasios"
// 	// tasos.lastName = "Theodosiou"
// 	// tasos.contact.email = "anastasios.theodosiou@gmail.com"
// 	// tasos.contact.zipCode = 54621

// 	tasos := person{
// 		firstName: "Anastasios",
// 		lastName:  "Theodoisou",
// 		contact: contactInfo{
// 			email:   "anastasios.theodosiou@gmail.com",
// 			zipCode: 54621,
// 		},
// 	}
// 	fmt.Println(tasos)
// 	// '%+v' list out every single property with its value
// 	fmt.Printf("%+v", tasos)
// }
