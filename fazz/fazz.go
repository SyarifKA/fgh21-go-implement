package fazz

import (
	"fazztrack/fazzfood/calc"
	"fmt"
)

func FazzFood(price int, promoCode string, inputDistance int, tax bool){
	deliveryFee := calc.CalcDeliveryFee(inputDistance)
	discount := calc.CalcDiscount(price, promoCode)
	taxValue := calc.CalcTax(price, tax)

	// Menampilkan harga
	fmt.Printf("Harga		: %d\n", price)
	fmt.Printf("Potongan	: %d\n", discount)
	fmt.Printf("Pajak		: %d\n", taxValue)
	fmt.Printf("Biaya Antar	: %d\n", deliveryFee)
	
	subTotal := (price + deliveryFee + taxValue)-discount
	fmt.Printf("SubTotal	: %d", subTotal)
}