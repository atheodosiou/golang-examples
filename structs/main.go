package main

import "fmt"

//Struct properties default values

// string => ""
// int    => 0
// float  => 0
// bool   => false
type person struct {
	firstName string
	lastName  string
}

func main() {
	fmt.Println("Working with structs!")
	//First type of declaration
	//This approach depends on the order we list out our properties...

	// tasos := person{"Anastasios", "Theodosiou"}
	// fmt.Println("Hello ", tasos.firstName+" "+tasos.lastName)

	//Second type of declaration

	// tasos := person{firstName: "Anastasios", lastName: "Theodosiou"}
	// fmt.Println("Hello ", tasos.firstName+" "+tasos.lastName)

	//Third way
	var tasos person
	tasos.firstName = "Anastasios"
	tasos.lastName = "Theodosiou"

	fmt.Println(tasos)
	// '%+v' list out every single property with its value
	fmt.Printf("%+v", tasos)
}
