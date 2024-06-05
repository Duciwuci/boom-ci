// Go program to illustrate the
// concept of main() function

// Declaration of the main package
package main

// Importing packages
import (
	"fmt"
)

// Main function
func main() {

	// Sorting the given slice
	faktoren := []int{345, 78}

	// Finding the time
	fmt.Println("Das Ergebnis ist: ", multipliziere(faktoren[0], faktoren[1]))

	// Finding the time
	fmt.Println("Das Ergebnis ist: ", multipliziere2(faktoren[0], faktoren[1]))

	// Finding the time
	fmt.Println("Das Ergebnis ist: ", multipliziere3(faktoren[0], faktoren[1]))

}
func multipliziere(faktor1, faktor2 int) int {
	return faktor1 * faktor2
}

func multipliziere2(faktor1, faktor2 int) int {
	return faktor1 * faktor2
}

func multipliziere3(faktor1, faktor2 int) int {
	return faktor1 * faktor2
}
