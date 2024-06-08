#include <cstdlib>
#include <iostream>
#include <functional>

using namespace std;

class Igrac {
public:
    string ime;
    int rejting;

    Igrac(string ime = "", int rejting = 0) : rejting(rejting), ime(ime) { }

    bool operator<(const Igrac &other) {
        if (rejting == other.rejting) {
            return ime < other.ime;
        }
        return rejting > other.rejting;
    }

};

class Opseg {
public:
    Opseg(int rejting = 1200) : rejting(rejting) {}

    bool operator()(const Igrac &i) {
        return abs(i.rejting - rejting) < 100;
    }
private:
    int rejting;
};

template<typename Kol>
int ispisi(const Kol &kolekcija, const function<bool(Igrac)> &f) {
    int brojac = 0;
    for (typename Kol::const_iterator it = kolekcija.cbegin(); it != kolekcija.cend(); it++) {
        if (f(*it)) {
            cout << it->ime << endl;
            brojac++;
        }
    }
    return brojac;
}



int main() {

    int n;
    cin >> n;
    vector<Igrac> niz(n);
    for(int i = 0; i < n; i++) {
        cin >> niz[i].ime >> niz[i].rejting;
    }

    int rejting;
    cin >> rejting;

    sort(niz.begin(), niz.end());
    for(int i = 0; i < n; i++) {
        cout << niz[i].ime << " " << niz[i].rejting << endl;
    }

    int rez = ispisi(niz, Opseg(rejting));

    if (rez == 0) {
        cout << "Nema odgovarajcih protivnika" << endl;
    }

    return 0;

}
