# prism
echelonのクライアント．以下の機能の実装を予定しています．

* Webカメラ制御
* 学生証リーダー
* ステータス表示
* メッセージボード
* WakeOnLan

詳しくは[TODO](TODO.md)を見てね．

機能はcameraとkioskにわかれてます．kioskもRasPiで動かそうと思っていましたが，
Qtのコードのコンパイルとかコンパイルとかコンパイルとかうまくいかないので諦めて
部室にあったネットブックにUbuntu入れてしのいでます．

camera，kioskどちらもGoで書かれています．kioskのUIはQMLで作りました．

## INSTALL camera on RasPi
追加で`camera/settings.go`が必要

[Unofficial ARM tarballs for Go](http://dave.cheney.net/unofficial-arm-tarballs)
からARMv6 multiarchのファイルをらずぱいにダウンロードし，`/usr/local/go`に展開
しておく．`$PATH`に`/usr/local/go/bin:/home/pi/go/bin`を追加し，`$GOPATH`を
`/home/pi/go`に設定

```bash
sudo apt-get install libopencv-dev
sudo vi /usr/lib/pkgconfig/opencv.pc
#`Libs: `の行に`-lm`を追記 (armではこれしないとうまくビルドできないぽい)
go get github.com/OUCC/prism/camera
cd $GOPATH/src/github.com/OUCC/prism/camera
cp /path/to/settings.go settings.go
go build -o camera
./camera

# autostart
sudo cp prism.init /etc/init.d/prism
sudo update-rc.d prism defaults
```

カメラは接続しておかないと起動しない．

## BUILD kiosk on Ubuntu 14.04
追加で`kiosk/settings.go`が必要

ビルド環境のアーキテクチャはデプロイ環境のと合わせる必要がある(i386, x86\_64)ので
注意

Ubuntu14.04のQt5.2ではQMLで日本語が豆腐になるバグがあったので，[Qt](qt.io)から
Linux Online Installerをダウンロードし，Qt5.4をインストールして対処した．Qt5.4が
デフォルトで入ってる場合は必要ない．以下Ubuntu14.04での作業．

goのセットアップは済ませたものとする．

```bash
export PKG_CONFIG_PATH=/opt/Qt/5.4/gcc/lib/pkgconfig
export LD_LIBRARY_PATH=/opt/Qt/5.4/gcc/lib
go get github.com/OUCC/prism/kiosk
cd $GOPATH/src/github.com/OUCC/prism/kiosk
cp /path/to/settings.go settings.go
go build -o kiosk
```

## INSTALL kiosk on Ubuntu 14.04
`kiosk`,`qml/`,`run.sh`をデプロイ環境にコピー．Qtを別途インストールした場合は
こちらにも同じバージョンをインストールしておく．

run.shは適宜編集すること．

```bash
sudo apt-get install fonts-migmix # QML内で使用
./run.sh
```

何回か実行しないと起動しなかったりする．カードリーダーは接続しないと起動しない．
起動したらカードを通してみて動くこと，日本語が表示されることを確認すること．

## LICENSE
MIT
