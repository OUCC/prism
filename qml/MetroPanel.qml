import QtQuick 2.2

Rectangle {
    property int _side: 45*unit
    property string size: 'medium'

    width: size === 'small'  ? _side : _side*2 + 5*unit
    height: size !== 'large' ? _side : _side*2 + 5*unit

    color: 'blue'
}
