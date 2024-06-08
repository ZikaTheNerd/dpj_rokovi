package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Stanica struct {
	id     int
	ime    string
	stanje int

	katanac sync.Mutex
}

type Autobus struct {
	id       int
	trenutna *Stanica

	posecenih int
	posecene  [3]bool
	gsp       *GSP
}

type GSP struct {
	autobusi [40]Autobus
	stanice  [3]Stanica
}

func prevezi(gsp *GSP, bus int, trenutna int, nova int) {
	trenutna_stanica := &gsp.stanice[trenutna]
	nova_stanica := &gsp.stanice[nova]
	autobus := &gsp.autobusi[bus]

	trenutna_stanica.katanac.Lock()
	fmt.Printf("Autobus broj %d napusta stanicu %s sa stanjem %d\n", bus, trenutna_stanica.ime, trenutna_stanica.stanje)
	trenutna_stanica.stanje--
	trenutna_stanica.katanac.Unlock()

	time.Sleep(time.Second * 2)

	nova_stanica.katanac.Lock()
	fmt.Printf("Autobus broj %d stize u stanicu %s sa stanjem %d\n", bus, nova_stanica.ime, nova_stanica.stanje)
	nova_stanica.stanje++

	if !autobus.posecene[nova] {
		autobus.posecenih++
		autobus.posecene[nova] = true
	}
	autobus.trenutna = nova_stanica
	nova_stanica.katanac.Unlock()
}

func vozi(bus *Autobus) {
	defer wg.Done()
	for {
		//fmt.Println(*bus)

		time.Sleep(time.Second * 2)
		nova_stanica := izaberi(bus.trenutna.id)
		prevezi(bus.gsp, bus.id, bus.trenutna.id, nova_stanica)

		if bus.posecenih == 3 {
			fmt.Printf("Autobus broj %d je obisao sve stanice i zavrsava voznju\n", bus.id)
			return
		}
	}
}

func izaberi(id int) int {
	for {
		nova_stanica := rand.Intn(3)
		if nova_stanica != id {
			return nova_stanica
		}
	}
}

func napraviGSP(gsp *GSP) {
	napraviStanice(gsp)
	napraviAutobuse(gsp)
}

func napraviStanice(gsp *GSP) {
	gsp.stanice[0] = Stanica{id: 0, ime: "Makedonska", stanje: 14}
	gsp.stanice[1] = Stanica{id: 1, ime: "Nemanjina", stanje: 10}
	gsp.stanice[2] = Stanica{id: 2, ime: "Vracar", stanje: 16}
}

func napraviAutobuse(gsp *GSP) {
	id := 0
	for stanica := 0; stanica < 3; stanica++ {
		for bus := 0; bus < gsp.stanice[stanica].stanje; bus++ {
			autobus := &gsp.autobusi[id]

			autobus.id = id
			autobus.trenutna = &gsp.stanice[stanica]
			autobus.gsp = gsp
			autobus.posecene = [3]bool{false, false, false}
			autobus.posecenih = 0
			id++
		}
	}
}

func (gsp *GSP) total() int {
	total := 0
	for stanica := 0; stanica < 3; stanica++ {
		total += gsp.stanice[stanica].stanje
	}
	return total
}

func main() {
	var gsp GSP
	napraviGSP(&gsp)
	wg.Add(40)
	for bus := 0; bus < 40; bus++ {
		go vozi(&gsp.autobusi[bus])
	}

	wg.Wait()

	fmt.Println("Svi busevi su obisli sve stanice bar jedanput")
	fmt.Println("Na kraju imamo", gsp.total(), "autobusa")
}
