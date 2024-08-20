package main

import (
	// "fazztrack/fazzfood/fazz"
	// "fazztrack/fazzfood/survey"

	"fazztrack/fazzfood/pointer"
)

type Bio struct{
	fullName string
	age int
	hobbies []string
}

func main(){
	// fazz.FazzFood(25000, "DITRAKTIR60", 5, true)

	// var hobbies [1]string //Arrays
	// hobbies[0] = "programming"
	// fmt.Println(hobbies) 

	// var hobbies []string = []string{ //Slice
	// 	"Soccer",
	// 	"Swimming",
	// 	"Anything",
	// 	"Coding",
	// 	"Etc",
	// }
	// fmt.Println(hobbies)

	// var bio Bio = Bio {
	// 	fullName: "fazztrack",
	// 	age: 20,
	// 	hobbies: []string{
	// 		"programming",
	// 		"Hiking",
	// 	},
	// }
	// fmt.Println(bio)

	// // survey.DataSurvey()
	// example.DataSurvey()
	// matrix.Matrix()

	pointer.TestPointer()
}