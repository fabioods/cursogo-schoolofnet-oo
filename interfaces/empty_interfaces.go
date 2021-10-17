package interfaces

import "fmt"

type Names []interface{}

func (n *Names) load() {
	*n = Names{
		"Javascript",
		"Java",
		"Python",
		"Ruby",
		"C++",
		"C#",
	}
}

func (n Names) print() {
	fmt.Println(n)
}

func ShowEmptyInterfaces() {
	fmt.Println("\n--- ShowEmptyInterfaces ---")
	var names Names
	names.load()
	names.print()

}
