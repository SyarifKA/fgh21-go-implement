package example

import (
	"fmt"
)
type Data struct{
	name string
}

func DataSurvey(){
	var skillset []Data = []Data{
		{
			name: "Javascript",
		},
		{
			name: "Semut",
		},
		{
			name: "PHP",
		},
		{
			name: "Gajah",
		},
	}
	fmt.Println(skillset[2].name)
	fmt.Println(skillset[0].name)
	// var Bio []Data = []Data {
	// 	{
	// 		skillset: []string {
	// 			"Sampoerna",
	// 		},
	// 	},
	// 	{
	// 		skillset: []string {
	// 			"Sampoerna",
	// 			"Esse",
	// 			"Cerutu",
	// 		},
	// 	},
	// 	{
	// 		skillset: []string {
	// 			"Sampoerna",
	// 			"Esse",
	// 			"Cerutu",
	// 		},
	// 	},
	// 	{
	// 		skillset: []string {
	// 			"Sampoerna",
	// 			"Esse",
	// 			"Cerutu",
	// 		},
	// 	},
	// }
	// for i := 0; i < len(Bio); i++{
	// 	fmt.Println(Bio[i])
	// }
	// fmt.Println(dataPerson)
}