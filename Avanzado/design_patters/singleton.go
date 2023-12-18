package main

import (
	"fmt"
	"sync"
	"time"
)

type Databse struct{}

var db *Databse
var lock sync.Mutex

func (Databse) CreateSingleConnection() {
	fmt.Println("Creating singlenton for Database")
	time.Sleep(2 * time.Second)
	fmt.Println("Creation Done")
}

func getDatabaseInstace() *Databse {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &Databse{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB connection already exists.")
	}
	return db
}

func mmain() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstace()
		}()

	}
	wg.Wait()
}
