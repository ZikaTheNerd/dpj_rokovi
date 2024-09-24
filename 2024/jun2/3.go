package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Gost struktura
type Gost struct {
	id           int
	trajanjeJela time.Duration
	mozeJesti    bool
	cond         *sync.Cond
}

// Restoran struktura
type Restoran struct {
	stolovi         int
	slobodniStolovi int
	mu              sync.Mutex
	condSto         *sync.Cond
	dostupniGosti   []*Gost
	condGost        *sync.Cond
	kraj            bool
}

func noviRestoran(brojStolova int) *Restoran {
	r := &Restoran{
		stolovi:         brojStolova,
		slobodniStolovi: brojStolova,
	}
	r.condSto = sync.NewCond(&r.mu)
	r.condGost = sync.NewCond(&r.mu)
	return r
}

func (r *Restoran) usluziGosta(g *Gost, wg *sync.WaitGroup) {
	defer wg.Done()

	r.mu.Lock()

	// Čekanje na slobodan sto
	for r.slobodniStolovi == 0 {
		fmt.Printf("Gost %d čeka na slobodan sto...\n", g.id)
		r.condSto.Wait()
	}

	// Gost zauzima sto
	r.slobodniStolovi--
	fmt.Printf("Gost %d je zauzeo sto. Slobodnih stolova: %d\n", g.id, r.slobodniStolovi)

	// Dodajemo gosta u red čekanja za konobara
	r.dostupniGosti = append(r.dostupniGosti, g)
	r.condGost.Signal() // Signalizujemo da je gost na redu da ga konobar usluži

	// Gost čeka da konobar završi uslugu pre nego što počne da jede
	for !g.mozeJesti {
		g.cond.Wait()
	}
	r.mu.Unlock()

	// Gost jede obrok (između 4 i 12 sekundi)
	fmt.Printf("Gost %d jede obrok...\n", g.id)
	time.Sleep(g.trajanjeJela)

	// Gost je završio i odlazi
	fmt.Printf("Gost %d je završio obrok i napušta restoran.\n", g.id)

	r.mu.Lock()
	r.slobodniStolovi++
	r.condSto.Signal() // Signalizacija da je sto slobodan
	r.mu.Unlock()
}

func (r *Restoran) konobarPosao(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		r.mu.Lock()

		// Ako nema gostiju koji čekaju i restoran nije zatvoren, konobar čeka
		for len(r.dostupniGosti) == 0 && !r.kraj {
			r.condGost.Wait()
		}

		// Ako je restoran zatvoren i nema više gostiju, konobar završava
		if r.kraj && len(r.dostupniGosti) == 0 {
			r.mu.Unlock()
			return
		}

		// Konobar preuzima prvog gosta iz reda
		gost := r.dostupniGosti[0]
		r.dostupniGosti = r.dostupniGosti[1:]

		fmt.Printf("Konobar %d uslužuje gosta %d.\n", id, gost.id)
		r.mu.Unlock()

		// Konobar uslužuje gosta (između 2 i 5 sekundi)
		vremeUsluge := time.Duration(rand.Intn(4)+2) * time.Second
		time.Sleep(vremeUsluge)

		// Nakon usluge, gost može početi jesti
		r.mu.Lock()
		gost.mozeJesti = true
		gost.cond.Signal()
		r.mu.Unlock()

		fmt.Printf("Konobar %d je završio sa usluživanjem gosta %d. (trajalo %v)\n", id, gost.id, vremeUsluge)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var brojKonobara int
	fmt.Print("Unesite broj konobara: ")
	fmt.Scan(&brojKonobara)

	restoran := noviRestoran(7)

	var wg sync.WaitGroup

	// Pokrećemo konobare kao gorutine
	for i := 1; i <= brojKonobara; i++ {
		wg.Add(1)
		go restoran.konobarPosao(i, &wg)
	}

	// Kreiramo 20 gostiju
	for i := 1; i <= 20; i++ {
		trajanjeJela := time.Duration(rand.Intn(9)+4) * time.Second
		gost := &Gost{id: i, trajanjeJela: trajanjeJela, mozeJesti: false}
		gost.cond = sync.NewCond(&restoran.mu)

		wg.Add(1)
		go restoran.usluziGosta(gost, &wg)

		// Nasumičan interval dolaska gostiju
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}

	// Čekamo da svi gosti završe
	wg.Wait()

	// Zaključujemo restoran i obaveštavamo konobare da je kraj
	restoran.mu.Lock()
	restoran.kraj = true
	restoran.condGost.Broadcast() // Obaveštavamo sve konobare da završe
	restoran.mu.Unlock()

	wg.Wait()
	fmt.Println("Svi gosti su usluženi. Restoran se zatvara.")
}
