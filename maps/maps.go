package maps

import "fmt"

func Show() {
	fmt.Println("\n\nUsando map")
	m := make(map[string]int)
	m["Fábio"] = 28
	m["Maria"] = 25
	m["Pedro"] = 23
	fmt.Println("Imprimindo o map: ", m)
	fmt.Println("Map na posiçao 'Fábio'", m["Fábio"])
	m["Fábio"] = 29
	fmt.Println("Atribui a posiçao 'Fábio' o valor 29", m)
	delete(m, "maria")
	fmt.Println("Tentei deletar do map a posiçao 'maria' que nao existe", m)
	delete(m, "Maria")
	fmt.Println("Deletei a posiçao 'Maria'", m)
	fmt.Println("Imprimi a posiçao 'Maria' que foi deletada", m["Maria"])
	_, exists := m["Maria"]
	fmt.Println("Verifica se existe a posiçao 'Maria'", exists)
	valuePedro, existsPedro := m["Pedro"]
	fmt.Println("Verifica se existe a posiçao 'Pedro'", existsPedro, valuePedro)

	fmt.Println("\nDeclarando map sem make e var")
	var mapWithoutMakeVar = map[string]int{}
	mapWithoutMakeVar["Fábio"] = 28
	mapWithoutMakeVar["Maria"] = 25
	mapWithoutMakeVar["Pedro"] = 23
	fmt.Println("Imprimindo o map sem make: ", mapWithoutMakeVar)

	fmt.Println("\nDeclarando map sem make")
	mapWithoutMake := map[string]int{"Altamir": 58, "Fábio": 28, "Maria": 25, "Pedro": 23}
	fmt.Println("Imprimindo o map sem make: ", mapWithoutMake)

	if value, exists := mapWithoutMake["Altamir"]; exists {
		fmt.Println("Existe o valor 'Altamir' no map", value)
	} else {
		fmt.Println("Não existe o valor 'Altamir' no map")
	}

}
