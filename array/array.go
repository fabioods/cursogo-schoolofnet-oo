package array

import "fmt"

func Show() {
	var x [10]int
	fmt.Println(x)
	fmt.Println(len(x))
	x[0] = 42
	x[2] = 99
	x[5] = -1
	fmt.Println(x)
	fmt.Println(x[0])
	// fmt.Println(x[10])
	var y [10]string
	fmt.Println(y)
	z := [3]int{1, 2, 3}
	fmt.Println(z)
}
