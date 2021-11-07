/*
Demo class for creational patterns
*/
package main

import "fmt"

func main() {
	fmt.Println("=========Factory example===========")
	var fruitFactory FruitFactorer = FruitFactory{}
	a := fruitFactory.getApple()
	b := fruitFactory.getBanana()
	o := fruitFactory.getOrange()
	fmt.Println(a, b, o)
	fmt.Println("=========Builder example===========")
	var addressBuilder *AddressBuilder = NewAddressBuilder()
	address := addressBuilder.withCountry("Russia").withCity("Saratov").withStreet("Moskovskya").withBuilding("10").withZip(410000).build()
	fmt.Println(address)
	fmt.Println("=========Singleton example===========")
	for i := 0; i < 10; i++ {
		getInstanse()
	}
	fmt.Println("=========Prototype example===========")
	var album Album = ReleaseAlbum(
		"Yellow Submarine / Eleanor Rigby",
		1966,
		[]string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"},
		[]string{"Eleanor Rigby", "Yellow Submarine"})

	disks := make([]*Album, 10)
	for i := 0; i < 10; i++ {
		al := album.clone()
		al.title += fmt.Sprint(i)
		disks[i] = al
	}
	for _, d := range disks {
		fmt.Println(d)
	}
}
