# NFC/Felica access sample for windows
Sony PaSoRiを使ったNFC/FelicaからGolangでIDM/PDMを取得するためのサンプル。

## Requirement
* PaSoRi Device
  * Sony RCS380でテスト。
* felica.dll
  * NFCポートソフトウェア(旧Felicaポートソフトウェア)に同梱されてます。
  　http://www.sony.co.jp/Products/felica/consumer/download/felicaportsoftware.html
* felicalib.dll
  * http://felicalib.tmurakam.org/
    コード内でDLL読み込みのためのパス指定が必要。
