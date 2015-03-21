import QtQuick 2.2

Text {
    id: loginMsg

    // 'login' or 'logout'
    property string info
    property string handleName
    property bool isFirstLogin

    text: info === 'in' ?
            isFirstLogin ? "はじめまして，" + handleName + "さん"
                : "おはようございます，" + handleName + "さん"
            : handleName + "氏，ログアウトしました"

    font.pixelSize: 16 * unit
    color: 'white'
}
