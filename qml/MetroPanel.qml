import QtQuick 2.2

Rectangle {
    property int _side: 180
    property string size: 'medium'

    width: size === 'small'  ? _side : _side*2 + 20
    height: size !== 'large' ? _side : _side*2 + 20

    color: 'blue'
}
