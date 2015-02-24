import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Window 2.1

ApplicationWindow {
    id: mainWindow
    title: qsTr("Prism")
    visibility: Window.FullScreen

    property int unit: 4

    Item {
        focus: true
        Keys.onEscapePressed: mainWindow.close()
    }

    FontLoader {
        id: fontawesome
        source: "font/fontawesome-webfont.ttf"
    }
    FontLoader {
        id: japaneseFont
        source: "/usr/share/fonts/truetype/migmix/migmix-1p-regular.ttf"
    }
    FontLoader {
        id: numberFont
        source: "/usr/share/fonts/truetype/liberation/LiberationSans-Regular.ttf"
    }

    BackgroundSwirls {
        anchors.fill: parent
    }

    Column {
        anchors.fill: parent
        anchors {
            topMargin: 12*unit; bottomMargin: 25*unit
            leftMargin: 25*unit; rightMargin: 25*unit
        }
        spacing: 12*unit

        Item {
            id: headerBar
            width: parent.width; height: indicator.height

            OnlineIndicator {
                id: indicator
                objectName: "indicator"
                anchors.right: parent.right
            }
        }

        MetroScreen {
            id: metro
        }
    }

    ReaderModal {
        id: readerModal
        anchors.fill: parent
        state: readerStatus.status === 'waiting' ? 'hide' : 'show'
    }
}
