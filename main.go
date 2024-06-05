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
	zahlen := []int{345, 79}

	// Finding the time
	fmt.Println("Ergebnis: ", multiplizieren(zahlen[0], zahlen[1]))

}
func multiplizieren(faktor1, faktor2 int) int {
	return faktor1 * faktor2
}
