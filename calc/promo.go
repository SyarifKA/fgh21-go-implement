package calc

func CalcDiscount(price int, promoCode string) int {
	// Perhitungan promo
	var discount int = 0
	if promoCode == "FAZZFOOD50" {
		if price >= 50000 {
			discount = price * 1 / 2
			if discount > 50000 {
				discount = 50000
			}
		}
	} else if promoCode == "DITRAKTIR60" {
		if price >= 25000 {
			discount = price * 6 / 10
			if discount > 30000 {
				discount = 30000
			}
		}
	}
	return discount
}

func CalcTax(price int, tax bool) int {
	var taxValue int = 0
	if tax == true {
		taxValue = price * 5 / 100
	}
	return taxValue
}

func CalcDeliveryFee(inputDistance int) int {
	var distance int = 2
	var distanceDifference int = 0
	var deliveryFee int = 5000
	if inputDistance < distance {
		deliveryFee = deliveryFee
	} else if inputDistance > distance {
		distanceDifference = inputDistance - distance
		deliveryFee = deliveryFee + (distanceDifference * 3000)
	}
	return deliveryFee
}