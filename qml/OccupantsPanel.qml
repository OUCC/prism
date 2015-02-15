import QtQuick 2.4

MetroPanel {
    size: 'large'
    color: '#DB532D'

    FontAwesomeIcon {
        anchors {
            left: parent.left
            leftMargin: 20
            bottom: parent.bottom
            bottomMargin: 20
        }

        size: 32
        icon: '\uf0c0'
    }

    Item {
        anchors {
            top: parent.top
            topMargin: 10
            bottom: parent.bottom
            bottomMargin: 10
            right: parent.right
            rightMargin: 20
        }

        Text {
            id: numberText

            anchors {
                top: parent.top
                right: peopleText.left
                rightMargin: 10
            }
            font.pixelSize: 64
            text: occupants.len
            color: 'white'
        }

        Text {
            id: peopleText

            anchors {
                bottom: numberText.bottom
                bottomMargin: 16
                right: parent.right
            }
            font.pixelSize: 24
            text: 'people'
            color: 'white'
        }

        Column {
            clip: true
            spacing: 5
            anchors {
                top: numberText.bottom
                topMargin: 10
                bottom: parent.bottom
                right: parent.right
            }

            Repeater {
                model: occupants.len
                delegate: Text {
                    text: occupants.get(index)
                    font.pixelSize: 24
                    color: 'white'
                }
            }
        }
    }
}
