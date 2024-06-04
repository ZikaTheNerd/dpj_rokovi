import matplotlib
matplotlib.use("QtAgg")

import json
import matplotlib.pyplot as plt
from PyQt5 import QtWidgets

class GUI(QtWidgets.QWidget):
    def __init__(self):
        super().__init__()
        self.setupUI()

    def setupUI(self):
        layout = self.setUpWindow(title="python", size=300)
        self.setupComponents(layout)
        self.setupConnections()

    def setUpWindow(self, title, size, spacing=20):
        self.setWindowTitle(title)
        width, height = size, size
        self.setGeometry(-width // 2, -height // 2, width, height)
        layout = QtWidgets.QGridLayout()
        layout.setSpacing(spacing)
        self.setLayout(layout)
        
        return layout

    def setupComponents(self, layout):
        self.label = QtWidgets.QLabel("Ime datoteke: ")
        self.lineEdit = QtWidgets.QLineEdit()
        self.lineEdit.setText("")
        self.button = QtWidgets.QPushButton("DIJAGRAM")
        layout.addWidget(self.label, 0, 0, 1, 3)
        layout.addWidget(self.lineEdit, 0, 4, 1, 3)
        layout.addWidget(self.button, 3, 0, 1, 3)
        

    def setupConnections(self):
        self.button.clicked.connect(self.drawDiagram)

    def drawDiagram(self):
        filename = self.lineEdit.text()
        list_days = []
        with open(filename, "r") as fp:
            list_days = json.load(fp)

        avg_temp = 0
        rain = 0
        day_map = dict()
        for day in list_days:
            day_name = day[0]
            indicator = day[1]
            temperature = day[2] 
            day_map[day_name] = temperature
            if indicator == "d":
                avg_temp += temperature
                rain += 1

        days = ["Ponedeljak", "Utorak", "Sreda", "Cetvrtak", "Petak", "Subota", "Nedelja"]
        plt.bar(days, list(map(day_map.get, days)))
        plt.title("Temperature u nedelji")
        plt.show()

        print(avg_temp / rain)

def main():
    app = QtWidgets.QApplication([])
    window = GUI()
    window.show()
    app.exec()

if __name__ == "__main__":
    main()
