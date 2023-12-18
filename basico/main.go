package main

import (
	"fmt"
)

/* HELLO WORLD
func main() {
	fmt.Println("Hello, World")
} */

/* func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.16

	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	//Declaracion de varaibles
	base := 12          //Primera forma
	var altura int = 14 //Segunda forma
	var area int        //Go no compila si las variables no son usadas

	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna vaalores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)
}
*/

/* func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
} */

func say(text string, c chan string) {
	c <- text
}

// func main() {
//clases
/* var myCar PK.CarPublic
myCar.Brand = "nissan"
myCar.Color = "Purple"
myCar.Owner = "Me"
myCar.Seating = 2
myCar.Year = 2023

fmt.Println("Los Datos de mi auto son:\n", myCar)

PK.PrintMessage("Hola desde main")

myCar.Ping()
myCar.DuplicateSeats()
fmt.Println(myCar) */

//------------------------------------------------------
//CONCURRENCIA GO ROUTINES
//Con Wait Group
/* var wg sync.WaitGroup
fmt.Println("Hello")
wg.Add(1)
go say("world", &wg)
wg.Wait()

go func(text string) {
	fmt.Println(text)
}("Adios")
time.Sleep(time.Second * 1) */

//Channels (organizar go routines)
/* c := make(chan string, 1)

fmt.Println("Hello")

go say("Bye", c)

fmt.Println(<-c) */
//}

func message(text string, c chan string) {
	c <- text
}
func main() {
	c := make(chan string, 2)

	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println(len(c), cap(c))

	// Range y close
	close(c)
	//c<-"Mensaje 3"

	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1", email1)
	go message("mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2)
		}
	}
}
