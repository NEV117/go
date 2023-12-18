package mypackage

import "fmt"

//CarPublic car con acceso publico
type CarPublic struct {
	Brand   string
	Year    int
	Seating int
	Color   string
	Owner   string
}

//poner las propiedades en minusculas lo hace privado
type CarPivate struct {
	brand   string
	year    int
	seating int
	color   string
	owner   string
}

//PrintMessage Publico
func PrintMessage(text string) {
	fmt.Println(text)
}

//PrintMessage Privado
func printMessage(text string) {
	fmt.Println(text)
}

//Ping Pong function
func (myClass CarPublic) Ping() {
	fmt.Println(myClass.Brand, "pong")
}

//Funcion de struc que modifica valor en memoria
func (myClass *CarPublic) DuplicateSeats() {
	myClass.Seating = myClass.Seating * 2
}

//to string
func (myClass CarPublic) String() string {
	return fmt.Sprintf("Brand: %s, Year: %d, Seating %d, Color: %s, Owner: %s", myClass.Brand, myClass.Year, myClass.Seating, myClass.Color, myClass.Owner)
}
