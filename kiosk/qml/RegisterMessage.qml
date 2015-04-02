import QtQuick 2.2

Column {
    id: registerMsg

    property alias felicaIDm : idmText.text
    property string info

    spacing: 10*unit

    Text {
        text: info === 'waiting'
                ? "FeliCaカードが登録されていません．
設定ページで以下のIDを入力し登録するか，
30秒以内にカードリーダーに学生証を通してください．"
                : info === 'success'
                ? "FeliCaカードを登録しました．"
                : "FeliCaカードの登録に失敗しました\n" + info
        font.pixelSize: 12*unit
        color: 'white'
    }
    Text {
        id: idmText
        visible: info === 'waiting'
        anchors.horizontalCenter: parent.horizontalCenter
        font.pixelSize: 18*unit
        color: 'white'
    }
}
