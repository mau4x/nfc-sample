package main

import (
	"log"
	"syscall"
	"unsafe"
)

func main() {
	var tPath = "D:/tools/felicalib-0.4.2/felicalib.dll"

	dll, err := syscall.LoadDLL(tPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dll.Release()

	// define procedures
	pasori_open, err := dll.FindProc("pasori_open")
	if err != nil {
		log.Fatal(err)
	}
	pasori_close, err := dll.FindProc("pasori_close")
	if err != nil {
		log.Fatal(err)
	}
	pasori_init, err := dll.FindProc("pasori_init")
	if err != nil {
		log.Fatal(err)
	}
	felica_polling, err := dll.FindProc("felica_polling")
	if err != nil {
		log.Fatal(err)
	}
	felica_free, err := dll.FindProc("felica_free")
	if err != nil {
		log.Fatal(err)
	}
	felica_getidm, err := dll.FindProc("felica_getidm")
	if err != nil {
		log.Fatal(err)
	}
	felica_getpmm, err := dll.FindProc("felica_getpmm")
	if err != nil {
		log.Fatal(err)
	}
//	felica_read_without_encryption02, err := dll.FindProc("felica_read_without_encryption02")
//	if err != nil {
//		log.Fatal(err)
//	}

	// errorチェックは、第一返り値の結果を判定
	// 第三返り値(lastErr)はいつもsuccess
	// lastErr != nilでチェックすると、successのときも値が入っている。
	tPasori, _, _ := pasori_open.Call(0)
	println("open", tPasori)
	if(tPasori == uintptr(0)) {
		log.Fatal("多分、端末が接続されてない")
	}
	defer pasori_close.Call(tPasori)

	ret, _, _ := pasori_init.Call(tPasori)
	println("init", ret)
	if(ret != 0) {
		log.Fatal("エラー出ることあるの？")
	}

	// tFelicaは、C側で内部的にメモリ確保されてる。
	// 従って多分開放しないと駄目
	tFelica, _, _ := felica_polling.Call(tPasori, 0xFFFF, 0, 0)
	println("poll", tFelica)
	if(tFelica == uintptr(0)) {
		log.Fatal("NFCカードの読み込み失敗。")
	}
	defer felica_free.Call(tFelica)

	// tFelicaで確保されえいる領域の一部を取得する。
	// 従って、失敗することはまずない。
	tIdm := uint64(1) // make([]byte, 8)
	felica_getidm.Call(tFelica, uintptr(unsafe.Pointer(&tIdm)))
	println("gIdm", tIdm)

	tPmm := uint64(1)
	felica_getpmm.Call(tFelica, uintptr(unsafe.Pointer(&tPmm)))
	println("gPmm", tPmm)
}