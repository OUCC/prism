import QtQuick 2.4

Text {
    property int size: 30
    property string icon

    font.family: fontawesome.name
    font.pixelSize: size
    text: icon
    color: 'white'
}
