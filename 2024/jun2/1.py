from PyQt5.QtWidgets import *
import json
import datetime

class MainWindow(QWidget):
    def __init__(self):
        super().__init__()
        self.setGeometry(200, 200, 300, 200)
        layout = QGridLayout()
        self.setLayout(layout)

        self.labelPocetak = QLabel("Pocetak intervala")
        self.labelKraj = QLabel("Kraj intervala")
        self.textBoxPocetak = QLineEdit("2024-06-02")
        self.textBoxKraj = QLineEdit("2024-06-05")
        self.labelPutanja = QLabel("Putanja do fajla")
        self.textBoxPutanja = QLineEdit("sales.json")
        self.buttonIzracunaj = QPushButton("Izracunaj")

        self.buttonIzracunaj.clicked.connect(self.on_button_clicked)

        layout.addWidget(self.labelPocetak, 1, 1, 1, 3)
        layout.addWidget(self.labelKraj, 1, 4, 1, 3)
        layout.addWidget(self.textBoxPocetak, 2, 1, 1, 3)
        layout.addWidget(self.textBoxKraj, 2, 4, 1, 3)
        layout.addWidget(self.labelPutanja, 3, 1, 1, 3)
        layout.addWidget(self.buttonIzracunaj, 3, 4, 1, 3)
        layout.addWidget(self.textBoxPutanja, 4, 1, 1, 3)

        self.setLayout(layout)

        # Create a placeholder for the result window
        self.result_window = None

    def on_button_clicked(self):
        self.calculate(self.textBoxPutanja.text(), self.textBoxPocetak.text(), self.textBoxKraj.text())

    def calculate(self, path, begin, end):
        # Create the result window
        self.result_window = QWidget()
        self.result_window.setGeometry(200, 200, 200, 100)
        resultLayout = QVBoxLayout()
        resultLabel = QLabel()
        resultLayout.addWidget(resultLabel)
        self.result_window.setLayout(resultLayout)

        try:
            with open(path, "r") as fp:
                obj = json.load(fp)
        except:
            resultLabel.setText("NEUSPESNO")
            self.result_window.show()
            return

        begin = datetime.datetime.strptime(begin, "%Y-%m-%d")
        end = datetime.datetime.strptime(end, "%Y-%m-%d")

        avg = 0
        max_price = -1
        maxProd = ""
        for product in obj:
            date = datetime.datetime.strptime(product["datum"], "%Y-%m-%d")
            if begin <= date <= end:
                price = product["cena"]
                if price >= max_price:
                    max_price = price
                    maxProd = product["ime artikla"]
                avg += price

        resultString = f"Ukupna zarada: {avg}\nNajskuplji artikal: {maxProd}"
        resultLabel.setText(resultString)
        self.result_window.show()


def main():
    app = QApplication([])
    main_window = MainWindow()
    main_window.show()
    app.exec()


if __name__ == "__main__":
    main()
