// Go program to illustrate the
// concept of main() function

// Declaration of the main package
package main

// Importing packages
import (
	"fmt"
)


func main() {
	zahlen := []int{345, 78}
	
	
	fmt.Println("Das Ergebnis der Multiplikation ist: ", multipliziere(zahlen[0],zahlen[1]))
	fmt.Println("Die Summe ist: ", addiere(zahlen[0],zahlen[1]))
	faktoren := []int{345, 78}

	
	fmt.Println("Das Ergebnis ist: ", multipliziere2(faktoren[0], faktoren[1]))

	
	fmt.Println("Das Ergebnis ist: ", multipliziere3(faktoren[0], faktoren[1]))

}
func multipliziere(faktor1, faktor2 int) int {
	return faktor1 * faktor2
}
func addiere(summand1,summand2 int) int {
	return summand1 + summand2 
}

func multipliziere2(faktor1, faktor2 int) int {
	return faktor1 * faktor2
}

func multipliziere3(faktor1, faktor2 int) int {
	return faktor1 * faktor2
}
