package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func names(values ...string) {
	for _, name := range values {
		fmt.Println(name)
	}
}
func getValues(x int) (double int, triple int, quad int) {
	// return 2 * x, 3 * x, 4 * x // vieja forma.
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println("sum", sum(1, 2, 4, 5, 6))
	names("a", "B", "c Name")

	// Imprimiento multiples returns:
	fmt.Println(getValues(5))

	// capturando multiples returns:
	d, t, q := getValues(5)
	fmt.Println(d, t, q)
}
