import matplotlib
matplotlib.use("Qt5Agg")

from PyQt5 import QtWidgets
import matplotlib.pyplot as plt
import numpy as np

class GUI(QtWidgets.QWidget):
    def __init__(self):
        super().__init__()
        self.initGUI()

    def initGUI(self):
        layout = self.setupWindow(size=300)
        self.setupComps(layout)
        self.setupConnections()

    def setupWindow(self, size, spacing=20):
        self.setWindowTitle("python")
        width, height = size, size
        self.setGeometry(-width // 2, -height // 2, width, height)
        layout = QtWidgets.QGridLayout()
        layout.setSpacing(spacing)
        self.setLayout(layout)

        return layout

    def setupComps(self, layout: QtWidgets.QGridLayout):
        self.label = QtWidgets.QLabel("Ime datoteke")
        self.lineEdit = QtWidgets.QLineEdit("anketa.txt")
        self.button = QtWidgets.QPushButton("DIJAGRAM")

        layout.addWidget(self.label, 0, 0, 1, 2)
        layout.addWidget(self.lineEdit, 0, 3, 1, 3)
        layout.addWidget(self.button, 3, 0, 1, 2)

    def setupConnections(self):
        self.button.clicked.connect(self.drawDiagram)

    def drawDiagram(self):
        filename = self.lineEdit.text()

        cityNames = []
        yes = []
        no = []
        with open(filename, "r") as fp:
            for line in fp.readlines():
                tokens = list(map(str.strip, line.split(",")))
                cityNames.append(tokens[0])
                yes.append(int(tokens[1]))
                no.append(int(tokens[2]))


        plt.bar(cityNames, yes, color="green", bottom=np.zeros(len(cityNames)), label="da")
        plt.bar(cityNames, no, color="red", bottom=np.zeros(len(cityNames)) + np.array(yes), label="ne")
        plt.ylabel("Broj osoba")
        plt.title("Stubicasti dijagram")
        plt.legend(loc="upper right")
        plt.show()

def main():
    app = QtWidgets.QApplication([])
    gui = GUI()
    gui.show()
    app.exec()

if __name__ == "__main__":
    main()
