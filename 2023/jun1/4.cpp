#include <iostream>
#include <istream>
#include <ostream>

template <typename T>
class Kvadar {
public:
    Kvadar() = default;

    Kvadar(T duzina, T sirina, T visina) : duzina{duzina}, visina{visina}, sirina{sirina} {
        if(duzina <= 0 || sirina <= 0 || visina <= 0) {
            throw "Nevalidne vrednosti";
        }
    }

    T zapremina() const {
        return duzina * sirina * visina;
    }

    bool operator >(const Kvadar &k) const {
        return zapremina() > k.zapremina();
    }
    bool operator ==(const Kvadar &k) const {
        return zapremina() == k.zapremina();
    }

    void ispis(std::ostream &os) const {
         os << "(" << duzina << ", " << sirina << ", " << visina << ")";
    }

    void upis(std::istream &is) {
        is >> duzina >> sirina >> visina;
    }

private:
    T duzina;
    T sirina;
    T visina;
    
};

template <typename T>
std::ostream &operator<<(std::ostream &os, const Kvadar<T> kvadar) {
    kvadar.ispis(os);
    return os;
}

template <typename T>
std::istream &operator>>(std::istream &is, Kvadar<T> kvadar) {
    kvadar.upis(is);
    return is;
}


int main() {
    double duzina1, sirina1, visina1;
    double duzina2, sirina2, visina2;

    std::cin >> duzina1 >> sirina1 >> visina1;
    std::cin >> duzina2 >> sirina2 >> visina2;

    Kvadar<double> k1;
    Kvadar<double> k2;
    try {
        k1 = Kvadar<double>(duzina1, sirina1, visina1);
        k2 = Kvadar<double>(duzina2, sirina2, visina2);
    } catch (const char *msg) {
        std::cout << msg << std::endl;
        return 0;
    }

    if(k1 > k2) {
        std::cout << k1 << std::endl;
        return 0;
    }
    if(k2 > k1) {
        std::cout << k2 << std::endl;
        return 0;
    }

}
