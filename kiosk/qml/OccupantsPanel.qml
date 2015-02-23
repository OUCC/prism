import QtQuick 2.2

MetroPanel {
    size: 'large'
    color: '#DB532D'

    FontAwesomeIcon {
        anchors {
            left: parent.left
            leftMargin: 5*unit
            bottom: parent.bottom
            bottomMargin: 5*unit
        }

        size: 8*unit
        icon: '\uf0c0'
    }

    Item {
        anchors {
            top: parent.top
            topMargin: 2*unit
            bottom: parent.bottom
            bottomMargin: 2*unit
            right: parent.right
            rightMargin: 5*unit
        }

        Text {
            id: numberText

            anchors {
                top: parent.top
                right: peopleText.left
                rightMargin: 2*unit
            }
            font.family: numberFont.name
            font.pixelSize: 16*unit
            text: occupants.len
            color: 'white'
        }

        Text {
            id: peopleText

            anchors {
                bottom: numberText.bottom
                bottomMargin: 4*unit
                right: parent.right
            }
            font.pixelSize: 6*unit
            text: 'people'
            color: 'white'
        }

        Column {
            clip: true
            spacing: 1*unit
            anchors {
                top: numberText.bottom
                topMargin: 2*unit
                bottom: parent.bottom
                right: parent.right
            }

            Repeater {
                model: occupants.len
                delegate: Text {
                    text: occupants.get(index)
                    font.pixelSize: 6*unit
                    color: 'white'
                }
            }
        }
    }
}
