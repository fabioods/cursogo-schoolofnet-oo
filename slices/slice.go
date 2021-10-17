package slices

import "fmt"

func Show() {
	fmt.Println()
	slice := make([]int, 2, 5)
	slice = append(slice, 1)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println()
	capSlice := make([]int, 2, 5) // tipo, tamanho, capacidade
	// capSlice[2] = 10
	// capSlice[3] = 20
	fmt.Println(capSlice)
	fmt.Printf("Cap %v \n", cap(capSlice))
	fmt.Printf("Lenght %v \n", len(capSlice))
	capSlice = append(capSlice, 7, 8, 9, 10)
	fmt.Println(capSlice)
	fmt.Printf("Cap %v \n", cap(capSlice))
	fmt.Printf("Lenght %v \n", len(capSlice))
	fmt.Println()

	for i := 0; i <= 20; i++ {
		capSlice = append(capSlice, i)
		fmt.Println("Cap: ", cap(capSlice), " Lenght: ", len(capSlice))
	}
	fmt.Println()
	sliceTest := slice
	slice[0] = 10
	slice = append(slice, 2, 3, 4)
	fmt.Println("slice test ", sliceTest, " cap ", cap(sliceTest))
	fmt.Println("slice ", slice, " cap ", cap(slice))

	fmt.Println("\n\nUsando slice sem make")
	sliceWithoutMake := []string{
		"Java", "javascript", "python", "go",
	}
	fmt.Println(sliceWithoutMake)
	sliceWithoutMake = append(sliceWithoutMake, "ruby")
	fmt.Println(sliceWithoutMake)
	fmt.Println(sliceWithoutMake[0])   // java
	fmt.Println(sliceWithoutMake[0:2]) // java javascript
	fmt.Println(sliceWithoutMake[1:2]) // javascript
	fmt.Println(sliceWithoutMake[2:])  // python go ruby
}
