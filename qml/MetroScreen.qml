import QtQuick 2.2

Row {
    spacing: 5*unit

    ClockPanel {
    }

    Column {
        spacing: parent.spacing

        OccupantsPanel {
        }

        MetroPanel {
            id: cardPanel
            size: 'small'
            color: '#AF193F'

            FontAwesomeIcon {
                anchors.centerIn: parent
                icon: '\uf09d'
                size: 32*unit
            }
        }
    }

    MetroPanel {
        id: messagePanel
        size: 'small'
        color: '#0089D1'

        FontAwesomeIcon {
            anchors.centerIn: parent
            icon: '\uf0e6'
            size: 32*unit
        }
    }
}
