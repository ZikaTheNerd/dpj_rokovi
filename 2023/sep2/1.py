import matplotlib
matplotlib.use("Qt5Agg")

from PyQt5 import QtWidgets
import matplotlib.pyplot as plt

class GUI(QtWidgets.QWidget):
    def __init__(self):
        super().__init__()

        self.createGUI()

    def createGUI(self):
        layout = self.setupWindow(size=300)
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
        self.label = QtWidgets.QLabel("Rezultati ispita")
        self.text = QtWidgets.QLineEdit("rezultati.txt")
        self.button = QtWidgets.QPushButton("DIJAGRAM")

        layout.addWidget(self.label, 0, 0, 1, 2)
        layout.addWidget(self.text, 0, 3, 1, 3)
        layout.addWidget(self.button, 1, 0, 1, 3)

    def setupConnections(self):
        self.button.clicked.connect(self.diagram)

    def diagram(self):
        filename = self.text.text()
        with open(filename, "r") as fp:
            grades = fp.readline().split(",")
            grades = map(str.strip, grades)
            grades = map(int, grades)
            grades = list(grades)

            plt.pie([grades[0], grades[1] + grades[2] + grades[3], grades[4] + grades[5]], labels=["Pali", "Srednji", "Najbolji"], colors=["red", "yellow", "green"])
            plt.title("Rezultati ispita")
            plt.show()
            

def main():
    app = QtWidgets.QApplication([])
    gui = GUI()
    gui.show()
    app.exec()

if __name__ == "__main__":
    main()
