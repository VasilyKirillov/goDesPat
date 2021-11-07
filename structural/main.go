/*
Demo class for structural patterns
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("================Adapter example=============")
	url := "https://www.boredapi.com/api/activity"
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal("can not get activity data")
		return
	}
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	activity := ActivityAdapter(result)
	fmt.Println(activity)

	fmt.Println("================Bridge example=============")
	u := User{id: 1, name: "Admin"}
	var ur1, ur2 BaseStorager
	ur1 = UserRepository{FileStorage{}}
	ur1.save(u)
	ur2 = UserRepository{DatabaseStorage{}}
	ur2.save(u)

	fmt.Println("================Proxy example=============")
	primes := PrimeNumberSequenceDecorator(1000)
	fmt.Println(primes)
	fmt.Println("================Facade example=============")
	caffee := Caffee{}
	caffee.getBreakfast()

	fmt.Println("================Flywaight example=============")
	pf := NewPenFactory()
	var redThinPen, blueThickPen, redThinPen2 Pen
	redThinPen = pf.GetThinPen(RED)
	redThinPen.draw("In vina veritas")
	blueThickPen = pf.GetThickPen(BLUE)
	blueThickPen.draw("To be or not to be")
	redThinPen2 = pf.GetThinPen(RED)
	redThinPen2.draw("Dolce vita")
}
