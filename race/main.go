package main

import (
	"fmt"
	"log"
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
		raceTime := 10000000000 / (rand.Intn(100) + 1)
		time.Sleep(time.Duration(raceTime) * time.Millisecond)
		raceT <- raceTime
		// wg2.Done()
	}()

	go func() {
		// timeout := time.Now().Add(1 * time.Minute)
		for {
			select {
			case <-ready:
				{
					fmt.Println("car is ready.")
					wg.Done()
				}
			case <-time.After(10 * time.Second):
				log.Fatal("car is stolen by UFO")
			}
		}
	}()
	go func() {
		// timeout := time.Now().Local().Add(1 * time.Minute)
		for {
			select {
			case x := <-raceT:
				{
					fmt.Printf("RaceTime %d .\n", x)
					wg2.Done()
				}
			case <-time.After(10 * time.Second):
				log.Fatal("car is stolen by UFO")
			}
		}
	}()
	wg.Wait()
	fmt.Println("START!!!")
	wg2.Wait()
}
