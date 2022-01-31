/* Instructions
In this exercise, you are going to write some code to help you prepare to buy a vehicle.

You have three tasks, one to determine if you need a license, one to help you choose between two vehicles and one to estimate the acceptable price for a used vehicle.

1. Determine if you will need a driver's license
Some vehicle kinds require a driver's license to operate them. Assume only the kinds "car" and "truck" require a license, everything else can be operated without a license.

Implement the NeedsLicense(kind) function that takes the kind of vehicle and returns a boolean indicating whether you need a license for that kind of vehicle.

needLicense := NeedsLicense("car")
// Output: true

needLicense = NeedsLicense("bike")
// Output: false
2. Choose between two potential vehicles to buy
You evaluate your options of available vehicles. You manage to narrow it down to two options but you need help making the final decision. For that, implement the function ChooseVehicle(option1, option2) that takes two vehicles as arguments and returns a decision that includes the option that comes first in dictionary order.

vehicle := ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
// Output: "Toyota Corolla is clearly the better choice."

ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf")
// Output: "Volkswagen Beetle is clearly the better choice."
3. Calculate an estimation for the price of a used vehicle
Now that you made a decision, you want to make sure you get a fair price at the dealership. Since you are interested in buying a used vehicle, the price depends on how old the vehicle is. For a rough estimate, assume if the vehicle is less than 3 years old, it costs 80% of the original price it had when it was brand new. If it is at least 10 years old, it costs 50%. If the vehicle is at least 3 years old but less than 10 years, it costs 70% of the original price.

Implement the CalculateResellPrice(originalPrice, age) function that applies this logic using if, else if and else (there are other ways if you want to practice). It takes the original price and the age of the vehicle as arguments and returns the estimated price in the dealership.

CalculateResellPrice(1000, 1)
// Output: 800

CalculateResellPrice(1000, 5)
// Output: 700

CalculateResellPrice(1000, 15)
// Output: 500
Note the return value is a float64. */

// package purchase
package main

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	//panic("NeedsLicense not implemented")
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	//panic("ChooseVehicle not implemented")
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	} else {
		return option2 + " is clearly the better choice."
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	//panic("CalculateResellPrice not implemented")

	if age >= 10 {
		return 0.7 * originalPrice
	} else if age >= 3 && age <= 10 {
		return 0.5 * originalPrice
	} else {
		return 0.8 * originalPrice
	}
}

func main() {

	needLicense := NeedsLicense("car")
	fmt.Println(needLicense)
	vehicle := ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
	fmt.Println(vehicle)
	vehicle = ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf")
	fmt.Println(vehicle)
	fmt.Println(CalculateResellPrice(1000, 1))
	fmt.Println(CalculateResellPrice(1000, 5))
	fmt.Println(CalculateResellPrice(1000, 15))
}
