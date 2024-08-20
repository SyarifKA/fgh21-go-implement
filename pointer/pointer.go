package pointer

import "fmt"

type Bio struct{
	name string
}

func TestPointer() {
	// a := 123
	// b := &a // & adalah referensi
	// *b = 345 // *b untuk menunjuk alamat yg sudah direferensikan
	// fmt.Println(a)
	// fmt.Println(*b)

	data := Bio{
		name: "fazz",
	}

	newData := &data.name
	*newData = "ilyas tampan"
	fmt.Println(data)
	fmt.Println(*newData)
}