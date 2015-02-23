import QtQuick 2.2

Row {
    property bool isOnline: true

    spacing: 4*unit

    Rectangle {
        id: circle
        anchors.verticalCenter: parent.verticalCenter
        width: 10*unit; height: width
        radius: width/2
        border {
            color: 'white'
            width: 1*unit
        }

        color: isOnline ? '#6CCB33' : 'red'
    }

    Text {
        id: txt

        width: txt_.width
        anchors.verticalCenter: parent.verticalCenter
        text: isOnline ? 'Online' : 'Offline'
        font.pixelSize: 10*unit
        color: 'white'
    }

    Text {
        id: txt_
        visible: false
        text: "Offline"
        font.pixelSize: txt.font.pixelSize
    }
}
