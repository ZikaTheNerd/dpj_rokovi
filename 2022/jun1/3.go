package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sastojci [4]string
var wg sync.WaitGroup

type Prodavnica struct {
	id        int
	kapacitet int
	stanje    int
	katanac   sync.Mutex
	nemaRobe  *sync.Cond
}

type Devojka struct {
	id                 int
	trenutnaProdavnica *Prodavnica
	sastojci           [4]int
	potrebno           [4]int
	centar             *Centar
	gotovo             int
}

type Centar struct {
	devojke    [5]Devojka
	prodavnice [4]Prodavnica
}

func napraviCentar(centar *Centar) {
	sastojci = [4]string{"Jagoda", "Sladoled", "Cokolada", "Banana"}
	napraviProdavnice(centar)
	napraviDevojke(centar)
}

func napraviProdavnice(centar *Centar) {
	for i := 0; i < 4; i++ {
		prodavnica := &centar.prodavnice[i]

		prodavnica.id = i
		prodavnica.kapacitet = rand.Intn(50001) + 50000
		prodavnica.stanje = prodavnica.kapacitet / 2
		prodavnica.nemaRobe = sync.NewCond(&prodavnica.katanac)

		go dopuna(centar, prodavnica)
	}
}

func napraviDevojke(centar *Centar) {
	for i := 0; i < 5; i++ {
		devojka := &centar.devojke[i]

		devojka.id = i
		devojka.sastojci = [4]int{0, 0, 0, 0}

		trenutna := rand.Intn(4)
		devojka.trenutnaProdavnica = &centar.prodavnice[trenutna]

		devojka.gotovo = 0
		for j := 0; j < 4; j++ {
			devojka.potrebno[j] = rand.Intn(3001)
		}
	}
}

func dopuna(centar *Centar, prodavnica *Prodavnica) {
	defer wg.Done()

	for {
		if prodavnica.kapacitet > prodavnica.stanje {
			kolicina := prodavnica.kapacitet - prodavnica.stanje
			dopuni(centar, prodavnica.id, kolicina)
		}

		time.Sleep(time.Second * 5)
	}
}

func dopuni(centar *Centar, id, kolicina int) {
	prodavnica := &centar.prodavnice[id]

	prodavnica.nemaRobe.L.Lock()
	prodavnica.stanje += kolicina
	prodavnica.nemaRobe.L.Unlock()
	prodavnica.nemaRobe.Broadcast()

}

func promeniProdavnicu(centar *Centar, devojka *Devojka) {
	trenutna := devojka.trenutnaProdavnica
	for {
		nova := &centar.prodavnice[rand.Intn(4)]
		if nova != trenutna {
			devojka.trenutnaProdavnica = nova
			break
		}
	}
}

func kupi(centar *Centar, devojka *Devojka) {
	defer wg.Done()
	for {
		trenutna_id := devojka.trenutnaProdavnica.id
		kolicina := devojka.potrebno[trenutna_id] - devojka.sastojci[trenutna_id]
		if kolicina > 0 {
			kupi2(centar, devojka.id, trenutna_id, kolicina)
		}

		if devojka.gotovo == 4 {
			fmt.Printf("Devojka broj %d je zavrsila kupovinu\n", devojka.id+1)
			return
		}

		time.Sleep(time.Second * 2)

		promeniProdavnicu(centar, devojka)

	}
}

func kupi2(centar *Centar, devojka_id int, prodavnica_id int, kolicina int) {
	devojka := &centar.devojke[devojka_id]
	prodavnica := &centar.prodavnice[prodavnica_id]

	prodavnica.nemaRobe.L.Lock()

	for prodavnica.stanje < kolicina {
		prodavnica.nemaRobe.Wait()
	}

	prodavnica.stanje -= kolicina
	devojka.sastojci[prodavnica_id] += kolicina
	devojka.gotovo++
	fmt.Printf("Devojka broj %d je kupila dovoljnu kolicinu sledeceg proizovda: %s\n", devojka_id+1, sastojci[prodavnica_id])

	prodavnica.nemaRobe.L.Unlock()
}

func main() {
	var centar Centar
	napraviCentar(&centar)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go kupi(&centar, &centar.devojke[i])
	}

	wg.Wait()
	fmt.Println("Sve devojke su kupile dovoljno sastojaka")
}
