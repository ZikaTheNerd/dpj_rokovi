#include <algorithm>
#include <cstdlib>
#include <iostream> 
#include <vector>

using namespace std;

class Protivnik {
public:
    string ime;
    unsigned int rejting;
    Protivnik() = default;
    Protivnik(const string &ime, unsigned int rejting) : ime(ime), rejting(rejting) {}

    bool operator<(const Protivnik &p) const {
        return rejting > p.rejting;
    }
};


template <typename Kol, typename F>
void ispisi(const Kol &k, const F& f) {
    bool flag = false;
    for (auto it = k.cbegin(); it != k.cend(); it++) {
        if (f(*it)) {
            flag = true;
            cout << it->ime << endl;
        }
    }
    if (!flag) {
        cout << "Nema odgovarajucih protivnika" << endl;
    }
}

class Opseg {
public:
    Opseg(unsigned int rejting = 1200) : rejting(rejting) {}

    bool operator ()(const Protivnik &p) const {
        return max(rejting, p.rejting) - min(rejting, p.rejting) <= 100;
    }

private:
    unsigned int rejting;
};

int main() {
    int n;
    cin >> n;
    vector<Protivnik> protivnici(n);
    for (int i = 0; i < n; i++) {
        cin >> protivnici[i].ime >> protivnici[i].rejting;
    }

    unsigned int rejting;
    cin >> rejting;

    sort(protivnici.begin(), protivnici.end());
    for (int i = 0; i < n; i++) {
        cout << protivnici[i].ime << " " << protivnici[i].rejting << endl;
    }

    Opseg f(rejting);
    ispisi(protivnici, f);

    return 0;
    
}
