import matplotlib
matplotlib.use("Qt5Agg")

from PyQt5 import QtWidgets
import matplotlib.pyplot as plt
import json

class GUI(QtWidgets.QWidget):
    def __init__(self):
        super().__init__()

        self.createGUI()

    def createGUI(self):
        layout = self.setupWindow(300)
        self.setupComps(layout)
        self.setupConnections()

    def setupWindow(self, size, spacing=20):
        width, height = size, size
        self.setGeometry(-width // 2, -height // 2, width, height)
        layout = QtWidgets.QGridLayout()
        layout.setSpacing(spacing)
        self.setLayout(layout)

        return layout

    def setupComps(self, layout):
        self.label = QtWidgets.QLabel("Ime fajla:")
        self.text = QtWidgets.QLineEdit("podaci.json")
        self.button = QtWidgets.QPushButton("DIJAGRAM")

        layout.addWidget(self.label, 0, 0, 1, 2)
        layout.addWidget(self.text, 0, 3, 1, 3)
        layout.addWidget(self.button, 1, 0, 1, 2)

    def setupConnections(self):
        self.button.clicked.connect(self.diagram)

    def diagram(self):
        filename = self.text.text()

        with open(filename, "r") as fp:
            data = json.load(fp)
        names = list(map(lambda person: person[0], data))
        heights = list(map(lambda person: person[3], data))
        weights = list(map(lambda person: person[2], data))
        boy_weights = list(map(lambda person: person[2], filter(lambda person: person[1] == "m", data)))
        girl_heights = list(map(lambda person: person[3], filter(lambda person: person[1] == "z", data)))

        print("Prosecna visina devojaka: ", sum(girl_heights) / len(girl_heights))
        print("Prosecna tezina decaka: ", sum(boy_weights) / len(boy_weights))

        plt.bar(names, heights, color="blue")
        plt.title("Visina dece u grupi")
        plt.show()

def main():
    app = QtWidgets.QApplication([])
    gui = GUI() 
    gui.show()
    app.exec()


if __name__ == "__main__":
    main()
