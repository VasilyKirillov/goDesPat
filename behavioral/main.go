package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=========Chain of responsibility pattern example===========")
	cashier := &cashier{}

	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "Duncan McClaud"}
	//Patient visiting
	reception.execute(patient)
	fmt.Println("=========Command pattern example===========")
	tv := &tv{}

	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()

	fmt.Println("=========Strategy pattern example===========")
	rand.Seed(time.Now().UTC().UnixNano())
	collection := make([]int, 20)
	for i := 0; i < 20; i++ {
		collection[i] = rand.Intn(100)
	}
	fmt.Println("before sort:", collection)
	sort(&collection, naturalOrder)
	fmt.Println("natural sort:", collection)
	sort(&collection, reverseOrder)
	fmt.Println("reverse sort:", collection)

	fmt.Println("=========Memento pattern example===========")
	te := TextEditor{TextWindow: TextWindow{}}
	te.write("Memento design pattern\n")
	te.write("implemented in Go\n")
	te.hitSave()

	te.write("incorrect")
	te.hitUndo()

	fmt.Println(te.print())

	fmt.Println("=========Observer pattern example===========")
	var magazine publisher = newMagazine("Murzilka")

	customer1 := &customer{id: "uncle Fedor"}
	customer2 := &customer{id: "Matroskin the cat"}

	magazine.subscribe(customer1)
	magazine.subscribe(customer2)

	magazine.notifyAll()
	magazine.unsubscribe(customer1)

}
