import QtQuick 2.4
import QtQuick.Controls 1.3
import QtQuick.Window 2.2

ApplicationWindow {
    id: mainWindow
    title: qsTr("Prism")
    visibility: Window.FullScreen
    minimumWidth: 800; minimumHeight: 640

    Item {
        focus: true
        Keys.onEscapePressed: mainWindow.close()
    }

    FontLoader {
        id: fontawesome
        source: "font/fontawesome-webfont.ttf"
    }

    BackgroundSwirls {
        anchors.fill: parent
    }

    Column {
        anchors.fill: parent
        anchors {
            topMargin: 50; bottomMargin: 100
            leftMargin: 100; rightMargin: 100
        }
        spacing: 50

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
