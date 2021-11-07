/*
Example of the Fruit factory
*/
package main

type Fruit struct{ name string }

func (f Fruit) String() string {
	return f.name
}

type Apple struct{ Fruit }
type Banana struct{ Fruit }
type Orange struct{ Fruit }

type FruitFactorer interface {
	getApple() Apple
	getBanana() Banana
	getOrange() Orange
}

type FruitFactory struct {
}

func (FruitFactory) getApple() Apple {
	return Apple{Fruit{"apple"}}
}
func (FruitFactory) getBanana() Banana {
	return Banana{Fruit{"banana"}}
}
func (FruitFactory) getOrange() Orange {
	return Orange{Fruit{"orange"}}
}
