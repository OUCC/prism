# prism
echelonのクライアント．以下の機能の実装を予定しています．

* Webカメラ制御
* 学生証リーダー
* ステータス表示
* メッセージボード
* WakeOnLan

詳しくは[TODO](TODO.md)を見てね．

## INSTALL on RasPi
`settings.go`が必要

1. [Beginner’s guide to cross-compile Qt5 on RaspberryPi](http://qt-project.org/wiki/RaspberryPi_Beginners_guide)に従いQt5をビルドしraspiにインストール．`/mnt/rasp-pi-rootfs/usr/local/qt5pi`を`scp`か何かでコピーすればよい．
2. raspiで`sudo cp /usr/local/qt5pi/lib/pkgconfig/* /usr/lib/pkgconfig/`を実行
3. `libopencv-dev libgles2-mesa-dev`をインストールしておく
4. [Unofficial ARM tarballs for Go](http://dave.cheney.net/unofficial-arm-tarballs)からARMv6 multiarchのファイルをらずぱいにダウンロードする．
5. `/usr/local/go`に展開，`$PATH`に`/usr/local/go/bin:/home/pi/go/bin`を追加
6. `$GOPATH`を`/home/pi/go`に設定
7. `go get github.com/OUCC/prism` 又はPCからソースコードをコピーする
8. `settings.go`をPCからコピーする
9. `go get github.com/OUCC/prism`
10. `LD_LIBRARY_PATH=/usr/local/qt5pi/lib prism`で起動

### go getでエラーが出る場合
#### github.com/gvalkov/golang-evdev/evdev
`github.com/gvalkov/golang-evdev/evdev/cdefs.go` の `EVIOCSCLOCKID = C.EVIOCSCLOCKID`をコメントする

#### github.com/lazywei/go-opencv
`/usr/lib/pkgconfig/opencv.pc`の`Libs: `の行に`-lm`を追記

## Develop
Go言語とQMLで実装しようと思ってます．

## LICENSE
MIT
