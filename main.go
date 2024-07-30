package main

import "fmt"

func fazzfood(price int, promoCode string, inputDistance int, tax bool){
	var deliveryFee int = 5000
	var distance int = 2
	var taxValue int = 0
	var discount int = 0
	var distanceDifference int = 0

	// Menampilkan harga
	fmt.Printf("Harga		: %d\n", price)
	
	// Perhitungan promo
	if promoCode == "FAZZFOOD50" {
		if price >= 50000{
			discount = price * 1/2
			if(discount > 50000){
				discount = 50000
			}
		}
	}else if promoCode == "DITRAKTIR60"{
		if(price >= 25000){
			discount = price * 6/10
			if discount > 30000 {
				discount = 30000
			}
		}
	}

	// Potongan
	fmt.Printf("Potongan	: %d\n", discount)
	
	// Perhitungan tax
	if tax == true {
		taxValue = price * 5/100
	}
	fmt.Printf("Pajak		: %d\n", taxValue)
	
	// Perhitungan biaya antar
	if inputDistance < distance {
		deliveryFee = deliveryFee
	}else if inputDistance > distance {
		distanceDifference = inputDistance - distance
		deliveryFee = deliveryFee + (distanceDifference*3000)
	}

	fmt.Printf("Biaya Antar	: %d\n", deliveryFee)
	
	var subTotal = (price + deliveryFee + taxValue)-discount
	fmt.Printf("SubTotal	: %d", subTotal)
}

func main(){
	fazzfood(25000, "DITRAKTIR60", 5, true)
}