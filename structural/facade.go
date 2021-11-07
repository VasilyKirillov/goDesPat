/*
Caffee is a facade for Cook and Waiter functions
*/
package main

import "fmt"

type Cook struct{}

func (c Cook) makeRice() {
	fmt.Println("cooking spicy rice...")
}

type Waiter struct{}

func (w Waiter) bringTea() {
	fmt.Println("making green tea with mint and lemon...")
}
func (w Waiter) serveBreakfast() {
	fmt.Println("serving the table...")
}

type Caffee struct {
	cook   Cook
	waiter Waiter
}

func (c Caffee) getBreakfast() {
	c.cook.makeRice()
	c.waiter.bringTea()
	c.waiter.serveBreakfast()
	fmt.Println("Breakfast is ready")
}
