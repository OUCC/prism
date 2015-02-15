import QtQuick 2.4

Modal {
    modalContent: Item {
        id: content
        anchors.fill: parent

        Column {
            anchors.centerIn: parent
            spacing: 40
            visible: readerStatus.status === 'posting'
            Row {
                spacing: 60
                FontAwesomeIcon {
                    icon: '\uf09d'
                    size: 80
                    rotation: 40
                }
                Text {
                    text: "Sending..."
                    font.pixelSize: 80
                    font.bold: true
                    color: 'white'
                }
            }
            Loading {
                anchors.horizontalCenter: parent.horizontalCenter
                size: 80
            }
        }

        Column {
            id: errorText
            anchors.centerIn: parent
            spacing: 20
            visible: readerStatus.status === 'error'
            Text {
                anchors.horizontalCenter: parent.horizontalCenter
                text: "Error"
                font.pixelSize: 80
                font.bold: true
                color: 'white'
            }
            Text {
                text: readerStatus.error
                font.pointSize: 24
                color: 'white'
            }
        }

        Text {
            id: loginText
            anchors.centerIn: parent
            text: readerStatus.data.firstLogin ? "はじめまして，" + readerStatus.data.handleName
                                                 + "さん" : readerStatus.data.handleName
                                                 + "氏 ログインしました"
            font.pixelSize: 64
            color: 'white'
            visible: readerStatus.data.event === 'in'
        }

        Text {
            id: logoutText
            anchors.centerIn: parent
            text: readerStatus.data.handleName + "氏 ログアウトしました"
            font.pixelSize: 64
            color: 'white'
            visible: readerStatus.data.event === 'out'
        }

        Text {
            id: timeText
            anchors {
                bottom: parent.bottom
                bottomMargin: 20
                right: parent.right
                rightMargin: 20
            }
            text: {
                var now = new Date()
                return (now.getMonth() + 1) + "月" + now.getDate(
                            ) + "日 " + now.getHours(
                            ) + "時" + now.getMinutes() + "分"
            }
            font.pixelSize: 32
            color: 'white'
            visible: readerStatus.status === 'posted'
        }
    }
}
