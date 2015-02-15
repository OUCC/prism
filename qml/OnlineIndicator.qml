import QtQuick 2.4

Row {
    property bool isOnline: true

    spacing: 15

    Rectangle {
        id: circle
        anchors.verticalCenter: parent.verticalCenter
        width: 40; height: width
        radius: width/2
        border {
            color: 'white'
            width: 3
        }

        color: isOnline ? '#6CCB33' : 'red'
    }

    Text {
        id: txt

        width: txt_.width
        anchors.verticalCenter: parent.verticalCenter
        text: isOnline ? 'Online' : 'Offline'
        font.pixelSize: 40
        color: 'white'
    }

    Text {
        id: txt_
        visible: false
        text: "Offline"
        font.pixelSize: txt.font.pixelSize
    }
}
