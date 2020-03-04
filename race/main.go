package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ready := make(chan bool, 2)
	timer := make(chan int, 2)
	raceT := make(chan int, 2)
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	wg.Add(2)
	wg2.Add(2)
	go func() {
		timeToStart := rand.Intn(1000) + 1
		time.Sleep(time.Duration(timeToStart) * time.Millisecond)
		ready <- true
		timer <- timeToStart
		// wg.Done()
		wg.Wait()
		raceTime := 100000 / (rand.Intn(100) + 1)
		time.Sleep(time.Duration(raceTime) * time.Millisecond)
		raceT <- raceTime
		// wg2.Done()
	}()
	go func() {
		timeToStart := rand.Intn(1000) + 1
		time.Sleep(time.Duration(timeToStart) * time.Millisecond)
		ready <- true
		timer <- timeToStart
		// wg.Done()
		wg.Wait()
		raceTime := 100000 / (rand.Intn(100) + 1)
		time.Sleep(time.Duration(raceTime) * time.Millisecond)
		raceT <- raceTime
		// wg2.Done()
	}()
	go func() {
		for {
			select {
			case _, ok := <-ready:
				if ok {
					fmt.Println("car is ready.")
					wg.Done()
				} else {
					fmt.Println("Channel closed!")
				}
			default:
			}
		}
	}()
	go func() {
		for {
			select {
			case x, ok := <-raceT:
				if ok {
					fmt.Printf("RaceTime %d .\n", x)
					wg2.Done()
				} else {
					fmt.Println("Channel closed!")
				}
			default:
			}
		}
	}()
	wg.Wait()
	fmt.Println("START!!!")
	wg2.Wait()
}
