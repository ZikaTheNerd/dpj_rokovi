import matplotlib
matplotlib.use("Qt5Agg") #OVO NIJE POTREBNO NA ISPITU

#ODAVDE POCINJE KOD
import pandas as pd
import matplotlib.pyplot as plt
import datetime
from  PyQt5.QtWidgets import *

class GUI(QWidget):
    def __init__(self):
        super().__init__()

        self.setGeometry(200, 200, 600, 300)

        layout = QGridLayout()
        self.label1 = QLabel("Putanja do fajla:")
        self.label2 = QLabel("Ime artikla:")
        self.text1 = QLineEdit("sales.csv")
        self.text2 = QLineEdit("Shoes")
        self.button = QPushButton("Prikazi Grafik")

        layout.addWidget(self.label1, 1, 1, 1, 3)
        layout.addWidget(self.text1, 1, 4, 1, 3)
        layout.addWidget(self.label2, 2, 1, 1, 3)
        layout.addWidget(self.text2, 2, 4, 1, 3)
        layout.addWidget(self.button, 3, 1, 1, 6)

        self.setLayout(layout)
        self.button.clicked.connect(self.diagram1)

    def diagram1(self):
        filename = self.text1.text()
        name = self.text2.text()

        data = pd.read_csv(filename)
        print(data)
        data = data[data["article"] == name] #filtrira
        data["date"] = pd.to_datetime(data["date"], format="%Y-%m-%d") #pretvorimo kolonu dates u datume po formatu format
        data = data.sort_values("date") #sortiramo tabelu po datumu
        print(data["date"])
        #pandas ima ugradjen plot (naravno moze i plt.plot(data["date"], data["sale"])
        #ovaj ugradjeni poziva matplotlib.pyplot
        data.plot("date", "sale")
        plt.xlabel("Datum")
        plt.ylabel("Prodaja")
        plt.title(f"Prodaja artikla {name}")
        plt.show()

    def diagram2(self):
        filename = self.text1.text()
        name = self.text2.text().strip()

        #ovako mozemo i sami da parsiramo csv
        #otovrimo fajl u editoru samo da vidimo koji je separator izmedju kolona
        #u ovom primeru koristi se ','
        #takodje treba preskociti prvu liniju u fajlu jer su to imena kolona

        fp = open(filename, "r") #za citanje
        fp.__next__() #preskace prvu liniju, moze naravno i sa if-om
        dates = []
        prices = []
        for line in fp:
            tokens = line.split(',')
            tokens = list(map(str.strip, tokens))
            if tokens[1] == name:
                dates.append(datetime.datetime.strptime(tokens[0],"%Y-%m-%d"))
                prices.append(int(tokens[2]))


        datesAndPrices = list(zip(dates, prices))
        datesAndPrices.sort()
        dates = list(map(lambda x: x[0], datesAndPrices))
        prices = list(map(lambda x: x[1], datesAndPrices))

        #BTW: Sortiranje nije neophodno da bi radilo za test primer koji je dat

        plt.plot(dates, prices)
        plt.xlabel("Datum")
        plt.ylabel("Prodaja")
        plt.title(f"Prodaja artikla {name}")
        plt.show()

        
        

def main(): 
    app = QApplication([])
    gui = GUI()
    gui.show()
    app.exec()

if __name__ == "__main__":
    main()
