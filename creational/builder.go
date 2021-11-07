/*
AddressBuilder contains methods that is used to build Address structure
*/
package main

import "fmt"

type Address struct {
	country  string
	city     string
	street   string
	building string
	zip      int
}

func (a Address) String() string {
	return fmt.Sprintf("%v, %v, %v, %v, %v", a.country, a.city, a.street, a.building, a.zip)
}

type AddressBuilder struct{ Address }

func NewAddressBuilder() *AddressBuilder {
	return &AddressBuilder{Address{}}
}

func (ab *AddressBuilder) withCountry(country string) *AddressBuilder {
	ab.country = country
	return ab
}
func (ab *AddressBuilder) withCity(city string) *AddressBuilder {
	ab.city = city
	return ab
}
func (ab *AddressBuilder) withStreet(street string) *AddressBuilder {
	ab.street = street
	return ab
}
func (ab *AddressBuilder) withBuilding(building string) *AddressBuilder {
	ab.building = building
	return ab
}
func (ab *AddressBuilder) withZip(zip int) *AddressBuilder {
	ab.zip = zip
	return ab
}
func (ab *AddressBuilder) build() *Address {
	return &ab.Address
}
