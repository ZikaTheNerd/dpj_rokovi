#include <iostream> 
#include <cmath> 

using namespace std;

template <typename T> 
class Elipsa {
public:
    Elipsa() = default;

    Elipsa(T a, T b) : a(a), b(b) {
        if (a <= 0 || b <= 0) {
            throw "Nevalidan unos";
        }
    }

    double povrsina() const {
        return a * b * M_PI * M_PI;
    }

    Elipsa operator+(const Elipsa &other) const {
        return Elipsa(a + other.a, b + other.b);
    }

    void ispis(ostream &os) const {
        os << "(" << a << ", " << b << ")";
    }
private:
    T a;
    T b;
};

int main() {
    double a1, b1;
    double a2, b2;

    cin >> a1 >> b1 >> a2 >> b2;

    Elipsa<double> e1;
    Elipsa<double> e2;
    try {
        e1 = Elipsa<double>(a1, b1);
        e2 = Elipsa<double>(a2, b2);
    } catch(const char *msg) {
        cout << msg << endl;
        return 0;
    }

    (e1 + e2).ispis(cout);

    return 0;
}
