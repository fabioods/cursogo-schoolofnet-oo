package interfaces

import "fmt"

type Vehicle interface {
	Start() string
}

func StartVehicle(v Vehicle) string {
	return v.Start()
}

type Car struct {
	Name string
}

type MotorCycle struct {
	Name string
}

func (c Car) Start() string {
	return fmt.Sprintf("The car %s has been started", c.Name)
}

func (m MotorCycle) Start() string {
	return fmt.Sprintf("The car %s has been started", m.Name)
}

func Show() {
	fmt.Println("\n\nInterfaces")

	car := Car{Name: "Ferrari"}
	fmt.Println(car.Start())
	fmt.Println("Start car with interface ", StartVehicle(car))

	motorcycle := MotorCycle{Name: "Honda"}
	fmt.Println(motorcycle.Start())
	fmt.Println("Start car with interface ", StartVehicle(motorcycle))
}
