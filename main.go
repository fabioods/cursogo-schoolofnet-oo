package main

import (
	"cursogo-schoolofnet-oo/array"
	"cursogo-schoolofnet-oo/interfaces"
	"cursogo-schoolofnet-oo/maps"
	"cursogo-schoolofnet-oo/slices"
	"cursogo-schoolofnet-oo/structs"
	"fmt"
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
	interfaces.ShowEmptyInterfaces()
}
