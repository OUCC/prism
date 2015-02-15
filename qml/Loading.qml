import QtQuick 2.4

Item {
    id: loading

    property alias size: circle.size

    width: circle.width; height: circle.height

    FontAwesomeIcon {
        id: circle
//        icon: '\uf1ce'
        icon: '\uf110'
        color: 'white'
    }

    RotationAnimation {
        running: loading.visible
        loops: Animation.Infinite
        target: circle
        from: 0; to: 360
        easing.type: Easing.InOutCubic
        duration: 1500
    }
}
