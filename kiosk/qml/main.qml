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

    // fonts
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

    // background animation
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
        objectName: "readerModal"
        anchors.fill: parent
    }

    FeliCaModal {
        id: felicaModal
        objectName: "felicaModal"
        anchors.fill: parent
    }

    // for test
//    Timer {
//        property int i: 0
//        interval: 1000
//        running: true
//        repeat: Animation.Infinite
//        onTriggered: {
//            if(i==0) readerModal.showReaderPosting();
//            if(i==1) readerModal.showReaderError("eee");
//            if(i==2) readerModal.showReaderInfo("login", "hoge", false);
//            i++;
//        }
//    }
}
