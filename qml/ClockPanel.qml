import QtQuick 2.2

MetroPanel {
    size: 'medium'
    color: '#5535B0'

    FontAwesomeIcon {
        anchors {
            left: parent.left; leftMargin: 20
            bottom: parent.bottom; bottomMargin: 20
        }

        size: 32
        icon: '\uf073'
    }

    Text {
        id: dateText

        anchors {
            top: parent.top; topMargin: 10
            right: parent.right; rightMargin: 40
        }
        font.family: numberFont.name
        font.pixelSize: 80
        color: 'white'
    }

    Text {
        id: dayText
        anchors {
            top: dateText.bottom; topMargin: -16
            horizontalCenter: dateText.horizontalCenter
        }
        font.pixelSize: 24
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
