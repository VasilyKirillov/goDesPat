package main

import "fmt"

//observable or subject or publisher

type publisher interface {
	subscribe(sub subscriber)
	unsubscribe(sub subscriber)
	notifyAll()
}

type magazine struct {
	subscriberList []subscriber
	name           string
	isPublished    bool
}

func newMagazine(name string) *magazine {
	return &magazine{
		name: name,
	}
}

func (i *magazine) subscribe(s subscriber) {
	i.subscriberList = append(i.subscriberList, s)
}

func (i *magazine) unsubscribe(s subscriber) {
	i.subscriberList = removeFromList(i.subscriberList, s)
}

func (i *magazine) notifyAll() {
	fmt.Printf("Magazine %s is published\n", i.name)
	i.isPublished = true
	for _, subscriber := range i.subscriberList {
		subscriber.update(i.name)
	}
}

func removeFromList(subscriberList []subscriber, subscriberToRemove subscriber) []subscriber {
	subscriberListLength := len(subscriberList)
	for i, subscriber := range subscriberList {
		if subscriberToRemove.getID() == subscriber.getID() {
			subscriberList[subscriberListLength-1], subscriberList[i] = subscriberList[i], subscriberList[subscriberListLength-1]
			return subscriberList[:subscriberListLength-1]
		}
	}
	return subscriberList
}

//subscriber subscriber
type subscriber interface {
	update(string)
	getID() string
}

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Customer %s is going to buy %s magazine\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}
