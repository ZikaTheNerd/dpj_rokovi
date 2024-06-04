package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	makedonska = "makedonska"
	nemanjina  = "nemanjina"
	vracar     = "vracar"
)

type Bus struct {
	id              int
	currentStation  string
	visitedStations map[string]bool
}

var (
	stations       = []string{makedonska, nemanjina, vracar}
	busesAtStation = map[string]int{
		makedonska: 14,
		nemanjina:  10,
		vracar:     16,
	}

	mu sync.Mutex
	wg sync.WaitGroup
)

func getNextStation(currentStation string) string {
	var newStation string
	for {
		newStation = stations[rand.Intn(len(stations))]
		if newStation != currentStation {
			break
		}
	}
	return newStation
}

func runBus(bus Bus) {
	defer wg.Done()
	for {
		time.Sleep(time.Second * 2)
		mu.Lock()
		busesAtStation[bus.currentStation]--
		fmt.Println("Bus", bus.id, "je napustio stanicu", bus.currentStation, "sa", busesAtStation[bus.currentStation], "autobusa.")
		mu.Unlock()

		nextStation := getNextStation(bus.currentStation)
		bus.visitedStations[nextStation] = true
		bus.currentStation = nextStation

		mu.Lock()
		busesAtStation[bus.currentStation]++
		fmt.Println("Bus", bus.id, "je stigao u stanicu", bus.currentStation, "sa", busesAtStation[bus.currentStation], "autobusa.")
		mu.Unlock()

		if len(stations) == len(bus.visitedStations) {
			fmt.Println("Bus", bus.id, "je obisao sve stanice")
			return
		}

		time.Sleep(time.Second * 2)
	}
}

func total() {
	fmt.Println("Konacna provera stanja: ")
	total := 0
	for _, station := range stations {
		fmt.Printf("%s: %d buseva\n", station, busesAtStation[station])
		total += busesAtStation[station]
	}
	fmt.Println("Ukupno:", total, "buseva")
}

func main() {
	busCount := 10 + 14 + 16
	for i := 0; i < busCount; i++ {
		var startStation string
		if i < 14 {
			startStation = makedonska
		} else if i < 24 {
			startStation = nemanjina
		} else {
			startStation = vracar
		}
		wg.Add(1)
		go runBus(Bus{id: i + 1, currentStation: startStation, visitedStations: map[string]bool{}})
	}

	wg.Wait()
	fmt.Println("Gotovo!")
	total()

}
