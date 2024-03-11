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
	root := "magisk"
	log.Println("Zygisk run on", runtime.GOARCH)
	zygiskv := os.Getenv("ZKSU_VERSION")
	log.Println("zygsik version", zygiskv)
	flag.Parse()
	args := flag.Args() //将传入的参数赋值给args变量
	if len(args) <= 0 {
		Start()
		return
	}
	if len(args) == 2 {
		return
	}
	if len(args) == 1 {
		switch args[0] {
		case "version":
			fmt.Println("zygsik version", zygiskv)
		case "root":
			fmt.Printf("root impl:%s", root)
		}

	}
	Loadmoudle(runtime.GOARCH)
}
