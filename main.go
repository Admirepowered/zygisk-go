package main

import (
	"log"
	"os"
	"runtime"
	"unsafe"
)

func is_64() bool {
	if unsafe.Sizeof(uintptr(0)) == 8 {
		return true
	} else if unsafe.Sizeof(uintptr(0)) == 4 {
		return false
	} else {
		log.Fatal("Unknow system arch")
		os.Exit(-9)
	}
	return false
}
func main() {
	log.Println("Zygisk run on", runtime.GOARCH)
	zygiskv := os.Getenv("ZKSU_VERSION")
	log.Println("zygsik version", zygiskv)
	Loadmoudle(runtime.GOARCH)
}
