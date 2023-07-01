package purchase

import (
	"strings"
)
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
        return true
    } 
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	solution := option2
	if strings.Compare(option1, option2) <= 0 {
		solution = option1
	}

	return solution + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	
	var price float64

	if age < 3 {
		price = originalPrice * 0.8
	}

	if age >= 3 && age < 10 {
		price = originalPrice * 0.7
	}

	if age >= 10 {
		price = originalPrice * 0.5
	}

	return price
}
