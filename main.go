package main

import (
	"fmt"
	"songo-oo/array"
	"songo-oo/interfaces"
	"songo-oo/maps"
	"songo-oo/slices"
	"songo-oo/structs"
)

func main() {
	fmt.Println("Array")
	array.Show()
	fmt.Println("Slices")
	slices.Show()
	maps.Show()
	structs.Show()
	fmt.Println("\nStructs sendo chamadas externas ao pacote")
	car := structs.Car{Name: "Ferrari"}
	fmt.Println("Carro externo", car.Info())
	structs.ShowHeranca()
	structs.ShowJSON()
	interfaces.Show()
}
