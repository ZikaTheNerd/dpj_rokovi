#include <iostream> 
#include <iterator>
#include <vector> 

using namespace std;

class Takmicar {
public:

    Takmicar() = default;

    Takmicar(unsigned int broj_takmicara, double poeni, string razred) :
        broj_takmicara(broj_takmicara), poeni(poeni), razred(razred) {}


    void upis(istream &is) {
        is >> broj_takmicara >> poeni >> razred;
    }

    void ispis(ostream &os) const {
        os << broj_takmicara << " " << poeni << " " << razred;
    }

private:
    unsigned int broj_takmicara;
    double poeni;
    string razred;
    friend class Plasman;

};

istream &operator>>(istream &is, Takmicar &t) {
    t.upis(is);

    return is;
}

ostream &operator<<(ostream &os, const Takmicar &t) {
    t.ispis(os);

    return os;
}

class Plasman {
public:
    Plasman(string razred, double granica):
        razred(razred), granica(granica) {}

    bool operator ()(const Takmicar &t) const {
        return (t.razred == razred && t.poeni >= granica);
    }
private:
    string razred;
    double granica;
};

template <typename Kol, typename Funkcional>
int transformisi(Kol &k, const Funkcional &f) {
    typename Kol::iterator it = k.begin();
    while (it != k.end()) {
        cout << "Sada sam na indeksu: " << it - k.begin() << endl;
        if (!f(*it)) {
            cout << "Brisem: " << *it << endl;
            k.erase(it);
        }
        else {
            it++;
        }
    }
    return k.size();
}

int main() {

    int n;
    cin >> n;
    vector<Takmicar> takmicari(n);

    for (int i = 0; i < n; i++) {
        cin >> takmicari[i];
    }

    for (const Takmicar &t: takmicari) {
        cout << t << endl;
    }

    string razred;
    double granica;
    cin >> razred >> granica;
    cout << transformisi(takmicari, Plasman(razred, granica)) << endl;

    for (const Takmicar &t: takmicari) {
        cout << t << endl;
    }

}

