import QtQuick 2.2

Column {
    id: errorMsg

    property alias errorText: text.text

    spacing: 5 * unit
    Text {
        anchors.horizontalCenter: parent.horizontalCenter
        text: "エラー"
        font.pixelSize: 20 * unit
        font.bold: true
        color: 'white'
    }
    Text {
        id: text
        font.pointSize: 6 * unit
        color: 'white'
    }
}
