package structs

import "fmt"

type SuperCar struct {
	Car
	CanFly bool
	Name   string
}

func (c SuperCar) InfoSuperCar() string {
	return fmt.Sprintf("%s, Info: %v, new Name: %s", c.Info(), c.CanFly, c.Name)
}

func ShowHeranca() {
	fmt.Println("\n\nHerança com struct")
	superCar := SuperCar{Car: Car{"Gol Z", 2012, "Vermelho"}, Name: "Gol y", CanFly: true}
	fmt.Println("superCar", superCar.InfoSuperCar())
}
