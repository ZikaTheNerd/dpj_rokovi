#include <iostream>

using namespace std;

template <typename T>
T euklid(T a, T b) {
    while (b != 0) { 
        int tmp = b;
        b = a % b;
        a = tmp;
    }

    return a;
}

template <typename T>
class Razlomak {
public:
    T brojilac;
    T imenilac;

    Razlomak(T brojilac, T imenilac) : brojilac(brojilac), imenilac(imenilac) {}

    void svedi() {
        T nzd = euklid(brojilac, imenilac);
        brojilac /= nzd;
        imenilac /= nzd;
    }

};

template <typename T>
ostream& operator <<(ostream &os, Razlomak<T> &r) {
    r.svedi();
    if (r.imenilac == 1 || r.brojilac == 0) {
        os << r.brojilac;
        return os;
    }

    os << r.brojilac << "/" << r.imenilac;
    return os;
}

int main() {
    int x, y;
    cin >> x >> y;
    Razlomak<int> r(x, y);
    cout << r << endl;

}
