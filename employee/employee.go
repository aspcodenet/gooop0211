package employee

import "fmt"

// struct vs class
// C#/Java class - en datastruktur som kan ha properties och metoder
// en Go struct  kan bara ha properties

// OOP - access modifiers (public, private, protected)
// OOP - constructors
// Dessa två hjälper oss att säkerställa VALID STATE
// tänk webbsida med formulär - vissa saker är MÅSTE MATA IN

type Employee struct {
	Name          string // mandatory
	Age           int    // mandatory
	StreetAddress string
	PostalCode    int
	City          string
}

// gentlemanna agreement
// I ett paket - finns det en metod som heter New - så är det en constructor
func New(name string, age int) Employee {
	return Employee{Name: name, Age: age}
}

func (e *Employee) PrintEmployee() { // detta blir en medlemsmetod för Employee
	fmt.Println(e.Name + " " + e.City)
}

// func PrintEmployee(e Employee) {
// 	fmt.Println(e.Name + " " + e.City)
// }

// skapa en constructor som tar mandatory fields
