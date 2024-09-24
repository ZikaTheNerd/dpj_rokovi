package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Centar struct {
	prijateljice [5]*Prijateljica
	prodavnice   [4]*Prodavnica
}

type Prijateljica struct {
	id       int
	potrebno [4]int
	korpa    [4]int
}

type Prodavnica struct {
	id        int
	kapacitet int
	stanje    int
	oznaka    byte

	katanac  sync.Mutex
	nemaRobe *sync.Cond
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func napraviCentar(c *Centar) {
	napraviPrijateljice(c)
	napraviProdavnice(c)
}

func napraviPrijateljice(c *Centar) {
	for i := 1; i <= 5; i++ {
		prijateljica := Prijateljica{id: i}
		for j := 0; j < 4; j++ {
			prijateljica.potrebno[j] = rand.Intn(1001)
			prijateljica.korpa[j] = 0
		}
		c.prijateljice[i-1] = &prijateljica
		wg.Add(1)
		go c.kupuj(i)
	}
}

func napraviProdavnice(c *Centar) {
	oznake := []byte{'J', 'S', 'C', 'B'}
	for i := 1; i <= 4; i++ {
		c.prodavnice[i-1] = &Prodavnica{id: i, kapacitet: 500, stanje: 100, oznaka: oznake[i-1]}
		c.prodavnice[i-1].nemaRobe = sync.NewCond(&c.prodavnice[i-1].katanac)
		go c.dopuni(i)
	}
}

func (c *Centar) dopuni(id int) {
	prodavnica := c.prodavnice[id-1]
	for {
		prodavnica.nemaRobe.L.Lock()
		prodavnica.stanje = prodavnica.kapacitet
		prodavnica.nemaRobe.L.Unlock()
		prodavnica.nemaRobe.Broadcast()
		fmt.Printf("Prodavnica %d se dopunila\n", id)
		time.Sleep(5 * time.Second)

	}
}

func (c *Centar) kupuj(id int) {
	defer wg.Done()
	posecene := []bool{false, false, false, false}
	brojPunih := 0
	trenutnaProdavnica := rand.Intn(4) + 1
	for brojPunih < 4 {
		for posecene[trenutnaProdavnica-1] {
			trenutnaProdavnica = rand.Intn(4) + 1
		}
		posecene[trenutnaProdavnica-1] = true

		c.kupi(id, trenutnaProdavnica)
		fmt.Printf("Prijateljica %d je kupila potrebnu kolicinu iz prodavnice %d\n", id, trenutnaProdavnica)
		brojPunih++
	}

}

func (c *Centar) kupi(idPrijateljice, idProdavnice int) {
	prijateljica := c.prijateljice[idPrijateljice-1]
	prodavnica := c.prodavnice[idProdavnice-1]

	razlika := prijateljica.potrebno[idProdavnice-1] - prijateljica.korpa[idProdavnice-1]
	prodavnica.nemaRobe.L.Lock()
	for razlika > 0 {

		for prodavnica.stanje == 0 {
			prodavnica.nemaRobe.Wait()
		}

		kolicina := min(razlika, prodavnica.stanje)
		prijateljica.korpa[idProdavnice-1] += kolicina
		razlika -= kolicina
		prodavnica.stanje -= kolicina
		fmt.Printf("Prijateljica %d je kupila %d %c iz prodavnice %d\n", idPrijateljice, kolicina, prodavnica.oznaka, idProdavnice)
	}
	prodavnica.nemaRobe.L.Unlock()

}

func main() {
	var c Centar
	napraviCentar(&c)
	wg.Wait()
	fmt.Println("GOTOVO")
}
