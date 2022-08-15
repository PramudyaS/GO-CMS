package main

import "fmt"

var chickens = []map[string]string{
	map[string]string{"name": "chicken blue", "gender": "male"},
	map[string]string{"name": "chicken red", "gender": "male"},
	map[string]string{"name": "chicken yellow", "gender": "female"},
}

func main() {
	for _, chicken := range chickens {
		printMessage(chicken["name"], chicken["gender"])
	}
}

func printMessage(name string, gender string) {
	var nameTxt, genderTxt = chikenMap(name, gender)
	fmt.Println(nameTxt, genderTxt)
}

func chikenMap(name string, gender string) (string, string) {
	nameTxt := name + "APPEND"

	genderTxt := gender + "APPEND"

	return nameTxt, genderTxt
}
