package main

import (
	"fmt"
	"sync"
	"time"
)

type Petospratnica struct {
	id           int
	gradjevinski int
	ostali       int

	gradjevinskiIskorisceno int
	ostaliIskorisceno       int
	gotovo                  bool

	katanac        sync.Mutex
	nemaMaterijala *sync.Cond
}

var petospratnice []Petospratnica
var wg sync.WaitGroup

func dopuniKamion(id int) {
	for {
		petospratnice[id].nemaMaterijala.L.Lock()
		if petospratnice[id].gotovo {
			petospratnice[id].nemaMaterijala.L.Unlock()
			break
		}
		petospratnice[id].gradjevinski += 10_000

		petospratnice[id].nemaMaterijala.L.Unlock()
		petospratnice[id].nemaMaterijala.Broadcast()
		time.Sleep(time.Second * 2)
	}
	wg.Done()
}

func dopuniKombi(id int) {
	for {
		petospratnice[id].nemaMaterijala.L.Lock()
		if petospratnice[id].gotovo {
			petospratnice[id].nemaMaterijala.L.Unlock()
			break
		}
		petospratnice[id].ostali += 5000

		petospratnice[id].nemaMaterijala.L.Unlock()
		petospratnice[id].nemaMaterijala.Broadcast()
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}

func iskoristi(id int) {
	for {
		petospratnice[id].nemaMaterijala.L.Lock()
		if petospratnice[id].gotovo {
			fmt.Println("Gradnja zgrade", id, "je gotova")
			petospratnice[id].nemaMaterijala.L.Unlock()
		}

		for petospratnice[id].gradjevinski < 100 || petospratnice[id].ostali < 50 {
			petospratnice[id].nemaMaterijala.Wait()
		}

		petospratnice[id].gradjevinski -= 100
		petospratnice[id].ostali -= 50
		petospratnice[id].gradjevinskiIskorisceno += 100
		petospratnice[id].ostaliIskorisceno += 50

		fmt.Println(id, ": Iskoriscen je materijal u zgradi")

		if petospratnice[id].gradjevinskiIskorisceno > 50_000 && petospratnice[id].ostaliIskorisceno > 50_000 {
			fmt.Println(id, ": Gotova izgradnja")
			petospratnice[id].gotovo = true
			petospratnice[id].nemaMaterijala.L.Unlock()
			break
		}
		petospratnice[id].nemaMaterijala.L.Unlock()
	}
	wg.Done()
}

func main() {

	petospratnice = make([]Petospratnica, 4)

	for id := 0; id < 4; id++ {
		petospratnice[id] = Petospratnica{id: id, gradjevinski: 0, ostali: 0,
			gradjevinskiIskorisceno: 0, ostaliIskorisceno: 0,
			gotovo: false}
		petospratnice[id].nemaMaterijala = sync.NewCond(&petospratnice[id].katanac)
	}

	wg.Add(12)

	for id := 0; id < 4; id++ {
		go dopuniKamion(id)
		go dopuniKombi(id)
		go iskoristi(id)
	}
	wg.Wait()

}
