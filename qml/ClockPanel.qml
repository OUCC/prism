import QtQuick 2.2

MetroPanel {
    size: 'medium'
    color: '#5535B0'

    FontAwesomeIcon {
        anchors {
            left: parent.left; leftMargin: 5*unit
            bottom: parent.bottom; bottomMargin: 5*unit
        }

        size: 8*unit
        icon: '\uf073'
    }

    Text {
        id: dateText

        anchors {
            top: parent.top; topMargin: 2*unit
            right: parent.right; rightMargin: 10*unit
        }
        font.family: numberFont.name
        font.pixelSize: 20*unit
        color: 'white'
    }

    Text {
        id: dayText
        anchors {
            top: dateText.bottom; topMargin: -4*unit
            horizontalCenter: dateText.horizontalCenter
        }
        font.pixelSize: 8*unit
        color: 'white'
    }

    Timer {
        running: true
        repeat: true
        interval: 1000
        onTriggered: {
            var now = new Date();
            var day_names = ["Sunday", "Monday", "Tuesday",
                "Wednesday", "Thursday", "Friday", "Saturday"];
            dateText.text = now.getDate();
            dayText.text = day_names[now.getDay()];
        }
    }
}
