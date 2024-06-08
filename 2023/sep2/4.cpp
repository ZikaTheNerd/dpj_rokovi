#include <iostream> 

using namespace std;

template<typename T>
class Kvadar {
private:
    T duzina;
    T sirina;
    T visina;
public:
    Kvadar() = default;


    Kvadar(T duzina, T sirina, T visina) : duzina(duzina), sirina(sirina), visina(visina) {
        if (duzina < 0 || sirina < 0 || visina < 0) {
            throw "Nevalidni podaci";
        }
    }

    T zapremina() const {
        return duzina * visina * sirina;
    }

    bool operator>(const Kvadar &other) const {
        return this->zapremina() > other.zapremina();
    }
    bool operator==(const Kvadar &other) const {
        return this->zapremina() == other.zapremina();
    }

    void ispis(ostream &os) const {
        os << "(" << duzina << ", " << sirina << ", " << visina << ")";
    }
};

template<typename T> 
ostream &operator<<(ostream &os, const Kvadar<T> &k) {
    k.ispis(os);
    return os;
}

int main() {

    double duzina1, sirina1, visina1;
    double duzina2, sirina2, visina2;
    Kvadar<double> k1;
    Kvadar<double> k2;

    cin >> duzina1 >> sirina1 >> visina1;
    cin >> duzina2 >> sirina2 >> visina2;

    try {
        k1 = Kvadar(duzina1, sirina1, visina1);
        k2 = Kvadar(duzina2, sirina2, visina2);
    } catch(const char *msg) {
        cout << msg << endl;
        return 0;
    }

    if(k1 > k2) {
        cout << k1 << endl;
    } else if (k2 > k1) {
        cout << k2 << endl;
    } else {
        cout << "Jednaki su" << endl;
    }

    return 0;
    
}
