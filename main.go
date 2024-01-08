package main

import (
	"cep/api/cep"
	"fmt"
)

func main() {
	cepReq := cep.RequestCepInfo("01001000")
	fmt.Println(cepReq)
	// prompt := promptui.Select{
	// 	Label: "Select Day",
	// 	Items: []string{
	// 		"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
	// 		"Saturday", "Sunday",
	// 	},
	// }
	//
	// _, result, err := prompt.Run()
	// fmt.Println(result)
	// fmt.Println(err)
}
