package structs

import (
	"encoding/json"
	"fmt"
)

func ShowJSON() {
	fmt.Println("\n\nUsando Json")
	car := Car{"Uno", 2021, "Black"}
	result, _ := json.Marshal(car)
	fmt.Println("Json Car", string(result))

	var carJson Car
	data := []byte(`{"name":"Uno","year":2020,"color":"Black"}`)
	json.Unmarshal(data, &carJson)
	fmt.Println("carJson", carJson)
}
