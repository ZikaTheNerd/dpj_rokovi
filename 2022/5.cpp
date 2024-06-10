#include <iostream>
#include <vector>

using namespace std;

class Cipela {
public:
    string model;
    string boja;
    double visina;

    Cipela() = default;
    Cipela(const string &model, const string &boja, double visina) : model(model), visina(visina), boja(boja) { }
};

istream &operator>>(istream &is, Cipela &c) {
    is >> c.model >> c.visina >> c.boja;
    return is;
}

class Odgovarajuce {
public:
    Odgovarajuce(double max_visina, const string &boja) : max_visina(max_visina), boja(boja) { }

    bool operator()(const Cipela &cipela) const {
        return cipela.boja == boja && cipela.visina <= max_visina;
    }
private:
    double max_visina;
    string boja;
};

template<typename Kol, typename F>
void filtriraj(Kol &k, const F& f) {
    typename Kol::iterator it = k.begin();
    while (it != k.end()) {
        if (!f(*it)) {
            k.erase(it);
            continue;
        }
        it++;
    }
}

int main() {
    int n;
    cin >> n;
    vector<Cipela> cipele(n);
    for (int i = 0; i < n; i++) {
        cin >> cipele[i];
    }

    string boja;
    double max_visina;
    cin >> max_visina >> boja;
    for (const auto &c : cipele) {
        cout << c.model << endl;
    }
    filtriraj(cipele, Odgovarajuce(max_visina, boja));

    if (cipele.size() == 0) {
        cout << "Nema odg cipela" << endl;
        return 0;
    }
    for (const auto &c : cipele) {
        cout << c.model << endl;
    }
    return 0;
}
