package main

import (
	"strings"

	"github.com/cheynewallace/tabby"
	"systementor.se/godemo0211/employee"
)

func createEmployee(name string, age int) *employee.Employee {
	e := employee.Employee{Name: name, Age: age} // Paxas/skapas en datastruktur 40 bytes
	return &e
}

// i andra språk NEW/malloc för att skapa variabler på HEAPEN
// i Go - använder ALDRIG NEW
// STACK OCH HEAPALLOKERING - det bestämmer kompilatorn inte VI UTVECKLARE

func main() {
	a := createEmployee("John", 30)
	a.PrintEmployee()
	a.StreetAddress = "dasdsadsaasdasddas"
	a.City = "Stockholm"
	name := "John"
	name = strings.ToLower(name)
	// john

	// oliver := employee.Employee{Name: "John", Age: 30, StreetAddress: "Testgatan 12", PostalCode: 0, City: ""}
	// fmt.Printf("%v är %v år", oliver.Name, oliver.Age)

	// //employee.PrintEmployee(oliver)
	// oliver.PrintEmployee()

	fanny := employee.New("Fanny", 25)
	fanny.PrintEmployee()

	tab := tabby.New()
	tab.AddHeader("NAME", "AGE", "STREET", "CITY")
	tab.AddLine(fanny.Name, fanny.Age, fanny.StreetAddress, fanny.City)
	tab.AddLine(a.Name, a.Age, a.StreetAddress, a.City)
	tab.Print()
	// josefine := employee.Employee{} // josefine är i INVALID STATE
	// fmt.Printf("%v är %v år", josefine.Name, josefine.Age)

}
