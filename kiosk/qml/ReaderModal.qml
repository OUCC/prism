import QtQuick 2.2

Modal {
    modalContent: Item {
        id: content
        anchors.fill: parent

        Column {
            anchors.centerIn: parent
            spacing: 10*unit
            visible: readerStatus.status === 'posting'
            Row {
                spacing: 15*unit
                FontAwesomeIcon {
                    icon: '\uf09d'
                    size: 20*unit
                    rotation: 10*unit
                }
                Text {
                    text: "Sending..."
                    font.pixelSize: 20*unit
                    font.bold: true
                    color: 'white'
                }
            }
            Loading {
                anchors.horizontalCenter: parent.horizontalCenter
                size: 20*unit
            }
        }

        Column {
            id: errorText
            anchors.centerIn: parent
            spacing: 5*unit
            visible: readerStatus.status === 'error'
            Text {
                anchors.horizontalCenter: parent.horizontalCenter
                text: "Error"
                font.pixelSize: 20*unit
                font.bold: true
                color: 'white'
            }
            Text {
                text: readerStatus.error
                font.pointSize: 6*unit
                color: 'white'
            }
        }

        Text {
            id: loginText
            anchors.centerIn: parent
            text: readerStatus.data.firstLogin ?
                      "Nice to meet you, " + readerStatus.data.handleName :
                      "Good morning, " + readerStatus.data.handleName
            font.pixelSize: 16*unit
            color: 'white'
            visible: readerStatus.data.event === 'in'
        }

        Text {
            id: logoutText
            anchors.centerIn: parent
            text: "See you, " + readerStatus.data.handleName
            font.pixelSize: 16*unit
            color: 'white'
            visible: readerStatus.data.event === 'out'
        }

        Text {
            id: timeText
            anchors {
                bottom: parent.bottom
                bottomMargin: 5*unit
                right: parent.right
                rightMargin: 5*unit
            }
            text: {
                var now = new Date()
                return (now.getMonth() + 1) + "/" + now.getDate(
                            ) + " " + now.getHours(
                            ) + ":" + now.getMinutes()
            }
            font.family: numberFont.name
            font.pixelSize: 8*unit
            color: 'white'
            visible: readerStatus.status === 'posted'
        }
    }
}
