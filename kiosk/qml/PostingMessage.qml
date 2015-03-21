import QtQuick 2.2

Column {
    id: postingMsg

    spacing: 10 * unit
    Row {
        spacing: 15 * unit
        FontAwesomeIcon {
            icon: '\uf09d'
            size: 20 * unit
            rotation: 40
        }
        Text {
            text: "送信中..."
            font.pixelSize: 20 * unit
            font.bold: true
            color: 'white'
        }
    }
    Loading {
        anchors.horizontalCenter: parent.horizontalCenter
        size: 20 * unit
    }
}
