package structs

import "fmt"

type SuperNumber int

type CarName string
type CarYear int
type CarColor string

type Car struct {
	Name  string `json:"name"`
	Year  int    `json:"-"`
	color string
}

func (c Car) Info() string {
	//return "Name: " + c.Name + " Year: " + string(c.Year) + " Color: " + c.color
	return fmt.Sprintf("Name: %s Year: %d Color: %s", c.Name, c.Year, c.color)
}

func (c Car) getName() string {
	return fmt.Sprintf("Car name is: %s - Info - %s", c.Name, c.Info())
}

func Show() {
	fmt.Println("\n\nUsando structs")
	var s SuperNumber = 10
	fmt.Println("Aplicando type a uma variável ", s)

	x := []SuperNumber{}
	x = append(x, 10, 20, 30)
	fmt.Println("Usando struct com slice ", x)

	var carName CarName = "Fusca"
	var carYear CarYear = 1965
	var carColor CarColor = "Azul"
	fmt.Println("Carro: ", carName, carYear, carColor)

	var fusca Car = Car{Name: "Fusca", Year: 1965, color: "Azul"}
	fmt.Println("Carro struct: ", fusca)

	gol := Car{Name: "Gol", Year: 2021, color: "Vermelho"}
	fmt.Println("Carro struct shorthand: ", gol)
	fmt.Println("Gol color: ", gol.color)
	fmt.Println("Gol info: ", gol.Info())
	fmt.Println("Gol name + info: ", gol.getName())

}
