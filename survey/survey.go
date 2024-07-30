package survey

import "fmt"

type Data struct{
	fullName string
	age int
	gender string
	isSmoker bool
	cigarVariant []string
}

func DataSurvey(){
	var dataPerson []Data = []Data {
		{
			fullName: "ilyas",
			age: 20,
			gender: "Pria",
			isSmoker: true,
			cigarVariant: []string {
				"Sampoerna",
				"Esse",
				"Cerutu",
			},
		},
		{
			fullName: "Fajry",
			age: 20,
			gender: "Pria",
			isSmoker: true,
			cigarVariant: []string {
				"Sampoerna",
				"Esse",
				"Cerutu",
			},
		},
		{
			fullName: "Fariq",
			age: 20,
			gender: "Pria",
			isSmoker: true,
			cigarVariant: []string {
				"Sampoerna",
				"Esse",
				"Cerutu",
			},
		},
	}
	for i := 0; i < len(dataPerson); i++{
		fmt.Println("Name ", dataPerson[i].fullName)
		fmt.Println("Age ", dataPerson[i].age)
		fmt.Println("gender ", dataPerson[i].gender)
		fmt.Println("isSmoker ", dataPerson[i].isSmoker)
		fmt.Println("cigarVariant ", dataPerson[i].cigarVariant)
	}
	// fmt.Println(dataPerson)
}