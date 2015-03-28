import QtQuick 2.2

Column {
    id: registerMsg

    property alias felicaIDm : idmText.text

    spacing: 10*unit

    Text {
        text: "FeliCaカードが登録されていません．\n設定ページで以下のIDを入力し登録してください．"
        font.pixelSize: 12*unit
        color: 'white'
    }
    Text {
        id: idmText
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: 18*unit
        color: 'white'
    }
}
