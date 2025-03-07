package main

import (
	"fmt"
)

func waitForDBs(numDBs int, dbChan chan struct{}) {

	for i := 0; i < numDBs; i++ {
		<-dbChan
	}

}

// don't touch below this line

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}
