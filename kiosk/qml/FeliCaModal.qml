import QtQuick 2.2
import QtMultimedia 5.0

Modal {
    id: felicaModal

    modalVisible: state !== 'waiting'
    modalColor: '#FF831E'

    function showFeliCaPosting() {
        state = 'posting';
        timer.stop();
    }

    function showFeliCaInfo(info, handleName, isFirstLogin) {
        state = 'info';
        infoMsg.info = info;
        infoMsg.handleName = handleName;
        infoMsg.isFirstLogin = isFirstLogin;

        timer.interval = 5000;
        timer.restart();
    }

    function showFeliCaRegistration() {
        state = 'register';
    }

    function showFeliCaError(text) {
        state = 'error';
        errorMsg.errorText = text;

        timer.interval = 10000;
        timer.restart();
    }

    state: 'waiting'

    states: [
        State {
            name: 'waiting'
        },
        State {
            name: 'posting'
        },
        State {
            name: 'info'
        },
        State {
            name: 'register'
        },
        State {
            name: 'error'
        }
    ]

    Timer {
        id: timer
        interval: 5000
        onTriggered: state = 'waiting'
    }

    modalContent: Item {
        id: content
        anchors.fill: parent

        PostingMessage {
            id: postingMsg

            anchors.centerIn: parent
            visible: felicaModal.state === 'posting'
        }

        ErrorMessage {
            id: errorMsg

            anchors.centerIn: parent
            visible: felicaModal.state === 'error'
        }

        InfoMessage {
            id: infoMsg

            anchors.centerIn: parent
            visible: felicaModal.state === 'info'
        }
    }
}

