/*
getInstance() method garantees that only one instance of the world is created
*/
package main

import (
	"fmt"
	"sync"
)

const (
	CREATING = "Creating the world"
	CREATED  = "World is already created"
)

type world struct{}

var oneWorld *world

var lock = &sync.Mutex{}

func getInstanse() *world {
	if oneWorld == nil {
		lock.Lock()
		defer lock.Unlock()
		if oneWorld == nil {
			fmt.Println(CREATING)
			oneWorld = &world{}
		} else {
			fmt.Println(CREATED)
		}
	} else {
		fmt.Println(CREATED)
	}
	return oneWorld
}
