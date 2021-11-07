/*
There are 2 types of storage: FileStorage and DatabaseStorage
also there is a Bridge interface that can work with one or another type of storage
*/
package main

import (
	"fmt"
	"time"
)

type User struct {
	id   int
	name string
}
type storager interface {
	store(u User) error
}

type FileStorage struct{}

func (fs FileStorage) store(u User) error {
	fmt.Println("Saving user data to the File")
	time.Sleep(500)
	return nil
}

type DatabaseStorage struct{}

func (dbs DatabaseStorage) store(u User) error {
	fmt.Println("Saving user data to the Database")
	time.Sleep(1000)
	return nil
}

//bridge interface
type BaseStorager interface {
	save(u User)
	getRepository() storager
}
type UserRepository struct {
	storager
}

func (ur UserRepository) save(u User) {
	ur.store(u)
}
func (ur UserRepository) getRepository() storager {
	return ur.storager
}
