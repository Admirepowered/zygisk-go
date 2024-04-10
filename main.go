package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"unsafe"
)

func Is_64() bool {
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
func Get_arch() string {
	cmd := "getprop ro.product.cpu.abi"

	// 执行命令
	output, err := exec.Command(cmd).Output()
	if err != nil {
		log.Println("command failed:", err)
		os.Exit(-8)
		return ""
	}
	return string(output)

}
func main() {
	zygiskv := "0.1"
	log.Println("Zygisk-go run on", runtime.GOARCH)
	log.Println("Zygsik-go version", zygiskv)
	flag.Parse()
	args := flag.Args() //将传入的参数赋值给args变量
	if len(args) <= 0 {
		Start()
		return
	}
	if len(args) > 1 {
		fmt.Println("Zygisk is runing in lib mod")
		return
	}

	Loadmoudle(runtime.GOARCH)
}
