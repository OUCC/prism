import QtQuick 2.2

Item {
    id: modal

    property alias modalContent: modalMain.children

    state: 'show'

    Rectangle {
        id: modalBack
        anchors.fill: parent
        color: 'black'
    }

    Rectangle {
        id: modalMain
        anchors.centerIn: parent
        width: parent.width
        height: parent.height / 2
        color: '#1E90FF'
    }

    states: [
        State {
            name: 'show'
            PropertyChanges {
                target: modalBack
                opacity: 0.5
            }
            PropertyChanges {
                target: modalMain
                opacity: 1
            }
        },
        State {
            name: 'hide'
            PropertyChanges {
                target: modalBack
                opacity: 0
            }
            PropertyChanges {
                target: modalMain
                opacity: 0
            }
        }
    ]

    transitions: Transition {
        NumberAnimation {
            target: modalBack; property: 'opacity'
            duration: 1000
        }
        NumberAnimation {
            target: modalMain; property: 'opacity'
            duration: 333
        }
    }
}
