#include <iostream>
#include <cmath>
#include <ostream>

namespace obrtna_tela {

template<typename T>
class PravaKupa {
public:
    PravaKupa() = default;

    PravaKupa(T h, T r) : h(h), r(r) {
        if (h < 0 || r < 0) {
            throw "-1";
        }
    }

    T zapremina() const {
        return M_PI * M_PI * h * r / 3.0;
    }

    bool operator<(const PravaKupa &other) const {
        return zapremina() < other.zapremina();
    }

    bool operator==(const PravaKupa &other) const {
        return zapremina() == other.zapremina();
    }

    bool operator!=(const PravaKupa &other) const {
        return zapremina() != other.zapremina();
    }

    void ispis(std::ostream &os) const {
        os << zapremina();
    }
private:
    T h;
    T r;
};


}

template <typename T>
std::ostream &operator<<(std::ostream &os, const obrtna_tela::PravaKupa<T> &pk) {
    pk.ispis(os);
    return os;
}

using namespace std;
using obrtna_tela::PravaKupa;

int main() {
    PravaKupa<double> pk1, pk2;

    double r1, h1, r2, h2;

    cin >> r1 >> h1 >> r2 >> h2;
    try {
        pk1 = PravaKupa(r1, h1);
        pk2 = PravaKupa(r2, h2);
    } catch (const char *msg) {
        cout << msg << endl;
        return 0;
    }

    if (pk1 != pk2) {
        cout << 0 << endl;
        cout << abs(pk1.zapremina() - pk2.zapremina()) << endl;
    } else {
        cout << 0 << endl;
        cout << pk1.zapremina() << endl;
    }
    return 0;
}
