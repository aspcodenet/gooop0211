package main

import (
	"fmt"

	"systementor.se/godemo0211/employee"
)

func createEmployee(name string, age int) *employee.Employee {
	e := employee.Employee{Name: name, Age: age} // Paxas/skapas en datastruktur 40 bytes
	return &e
}

// i andra språk NEW/malloc för att skapa variabler på HEAPEN
// i Go - använder ALDRIG NEW
// STACK OCH HEAPALLOKERING - det bestämmer kompilatorn inte VI UTVECKLARE

type Player struct {
	Name         string // Peter Forsberg
	Age          int    // 51
	JerseyNumber int    // 21
	Team         string // colorado
}

// List<Player> players = new List<Player>();
// Array<Player> players = new Array<Player>();

func modifyPlayer(foppa *Player) { // COPY BY VALUE
	foppa.Name = "BLABLA"
	fmt.Println(foppa.Name)
}

// void modifyPlayer(Player foppa) {  // COPY BY VALUE
// 	foppa.Name = "BLABLA";
// 	fmt.Println(foppa.Name);
// }

func main() {
	//    var foppa = new Player("Peter Forsberg");
	//    modifyPlayer(foppa)
	//    fmt.Println(foppa.Name)  // SKRIVER UT BLABLA

	foppa := Player{Name: "Peter Forsberg", Age: 51, JerseyNumber: 21, Team: "Colorado"}
	modifyPlayer(&foppa)
	fmt.Println(foppa.Name)

	// zata := Player{Name: "Henrik Zetterberg", Age: 44, JerseyNumber: 40, Team: "Detroit"}
	// zata2 := Player{Name: "2", Age: 2341, JerseyNumber: 2234, Team: "Colorado"}
	// zata3 := Player{Name: "2", Age: 2341, JerseyNumber: 2234, Team: "Colorado"}

	// var players = []Player{zata, zata2, foppa, zata3} // tommar [] ÄR DET EN SLICE (List<T>)
	// // Konsekutivt minne (alla saker ligger efter varandra i minnet)
	// // DYNAMISK (går inte att lägga till en tlil)

	// stefan := Player{Name: "Stefan", Age: 2341, JerseyNumber: 2234, Team: "Colorado"}
	// // I minnet zata, zata2, foppa, zata3 -
	// // kanske inte finns plats för zata3

	// players = append(players, stefan)
	// // 1 funktionsorienterat std lib
	// // 2 C - realloc för att öka en array

	// // vet hur många
	// players[0].Name = "Stefan"
	// for _, player := range players {
	// 	fmt.Println(player.Name)
	// }

	// // SSoT - Single Source of Truth
	// // var players = [...]Player{zata, zata2, foppa, zata3} // ARRAY - Konsekutivt minne (alla saker ligger efter varandra i minnet)
	// // // 3 * 40 = 120 bytes                       // STATISK (går inte att lägga till en tlil)
	// // // vet hur många
	// // players[0].Name = "Stefan"
	// // for _, player := range players {
	// // 	fmt.Println(player.Name)
	// // }

	// // for i := 0; i < len(players); i++ {
	// //     fmt.Println(players[i].Name)
	// // }

	// // bok medf 3 st sidor - 0, 1, 2
	// //                        0, 100, 7000

	// a := createEmployee("John", 30)
	// a.PrintEmployee()
	// a.StreetAddress = "dasdsadsaasdasddas"
	// a.City = "Stockholm"
	// name := "John"
	// name = strings.ToLower(name)
	// // john

	// // oliver := employee.Employee{Name: "John", Age: 30, StreetAddress: "Testgatan 12", PostalCode: 0, City: ""}
	// // fmt.Printf("%v är %v år", oliver.Name, oliver.Age)

	// // //employee.PrintEmployee(oliver)
	// // oliver.PrintEmployee()

	// fanny := employee.New("Fanny", 25)
	// fanny.PrintEmployee()

	// tab := tabby.New()
	// tab.AddHeader("NAME", "AGE", "STREET", "CITY")
	// tab.AddLine(fanny.Name, fanny.Age, fanny.StreetAddress, fanny.City)
	// tab.AddLine(a.Name, a.Age, a.StreetAddress, a.City)
	// tab.Print()
	// // josefine := employee.Employee{} // josefine är i INVALID STATE
	// // fmt.Printf("%v är %v år", josefine.Name, josefine.Age)

}
