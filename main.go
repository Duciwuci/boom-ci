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
	liste := []int{345, 78}
	
	
	// Finding the time
	fmt.Println("ergebnis: ", multipliziere(liste[0],liste[1]))

}
func multipliziere(faktor1,faktor2 int) int {
	return faktor1 * faktor2
}