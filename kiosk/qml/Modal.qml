import QtQuick 2.2

Item {
    id: modal

    property alias modalContent: modalMain.children
    property alias modalColor: modalMain.color
    property bool modalVisible

    Rectangle {
        id: modalBack
        anchors.fill: parent
        color: 'black'
        opacity: 0
    }

    Rectangle {
        id: modalMain
        anchors.centerIn: parent
        width: parent.width
        height: parent.height / 2
        opacity: 0
    }

    onModalVisibleChanged: {
        if (modalVisible) {
            inAnim.restart();
        }
        else {
            outAnim.restart();
        }
    }

    ParallelAnimation {
        id: inAnim
        OpacityAnimator {
            target: modalBack
            duration: 500
            from: 0; to: 0.5
        }
        OpacityAnimator {
            target: modalMain
            duration: 500
            from: 0; to: 1
        }
    }

    ParallelAnimation {
        id: outAnim
        OpacityAnimator {
            target: modalBack
            duration: 500
            from: 0.5; to: 0
        }
        OpacityAnimator {
            target: modalMain
            duration: 500
            from: 1; to: 0
        }
    }
}
