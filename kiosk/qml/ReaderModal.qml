import QtQuick 2.2

Modal {
    id: readerModal

    modalVisible: state !== 'waiting'
    modalColor: '#1E90FF'

    function showReaderPosting() {
        state = 'posting';
        timer.stop();
    }

    function showReaderInfo(info, handleName, isFirstLogin) {
        state = 'info';
        infoMsg.info = info;
        infoMsg.handleName = handleName;
        infoMsg.isFirstLogin = isFirstLogin;

        timer.interval = 5000;
        timer.restart()
    }

    function showReaderError(text) {
        state = 'error';
        errorMsg.errorText = text;

        timer.interval = 10000;
        timer.restart()
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
            visible: readerModal.state === 'posting'
        }

        ErrorMessage {
            id: errorMsg

            anchors.centerIn: parent
            visible: readerModal.state === 'error'
        }

        InfoMessage {
            id: infoMsg

            anchors.centerIn: parent
            visible: readerModal.state === 'info'
        }
    }
}
