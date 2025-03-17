package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = sync.RWMutex{}

// var mutex = sync.Mutex{}
var waitGroup = sync.WaitGroup{}
var dbData = []string{"1", "2", "3", "4", "5"}
var results = []string{}

func main() {
	var now time.Time = time.Now()
	for i := 0; i < len(dbData); i++ {
		waitGroup.Add(1)
		go dbCall(i)
	}
	waitGroup.Wait()
	fmt.Printf("\nTotal execution time: %v\n", time.Since(now))
	fmt.Printf("\nResults: %v\n", results)

}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Microsecond)
	save(dbData[i])
	log()
	waitGroup.Done()
}

func save(result string) {
	mutex.Lock()
	results = append(results, result)
	mutex.Unlock()
}

func log() {
	mutex.RLock()
	fmt.Printf("\nCurrent results: %v", results)
	mutex.RUnlock()
}
