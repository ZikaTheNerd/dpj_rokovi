#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Igracka {
public:
    string naziv;
    unsigned int cena;
    unsigned int godinaProizvodnje;

    Igracka(const string &naziv = "", unsigned int cena = 0, unsigned int godinaProizvodnje = 0) :
        naziv(naziv), cena(cena), godinaProizvodnje(godinaProizvodnje) {}

    bool operator <(const Igracka &druga) const {
        if (godinaProizvodnje == druga.godinaProizvodnje) {
            return cena > druga.cena;
        }

        return godinaProizvodnje < druga.godinaProizvodnje;
    }
};

class NajboljaIgracka {
public:
    unsigned int novac;
    Igracka najboljaDoSad;

    NajboljaIgracka(unsigned int novac, unsigned int cena = 0, unsigned int godinaProizvodnje = 0) : 
        novac(novac), najboljaDoSad("", cena, godinaProizvodnje) {}

    bool operator()(const Igracka &i) {
        if (novac < i.cena) {
            return false;
        }

        if (najboljaDoSad < i) {
            najboljaDoSad = i;
            return true;
        }

        return false;
    }
};

template<typename Kol>
Igracka odrediNajbolju(Kol &k, unsigned int novac) {
    NajboljaIgracka najbolja(novac);
    for (auto it = k.begin(); it != k.end(); it++) {
        najbolja(*it);
    }

    return najbolja.najboljaDoSad;
}

int main() {
    int n;
    cin >> n;
    vector<Igracka> igracke(n);
    for (int i = 0; i < n; i++) {
        string naziv;
        unsigned int cena, godinaProizvodnje;
        cin >> naziv >> godinaProizvodnje >> cena;

        igracke[i] = Igracka(naziv, cena, godinaProizvodnje);
    }

    unsigned int novac;
    cin >> novac;
    Igracka rez = odrediNajbolju(igracke, novac);



    cout << (rez.naziv == "" ? "Nema adekvatne" : rez.naziv) << endl;

}
