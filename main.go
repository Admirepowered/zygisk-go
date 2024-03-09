package main

import (
	"fmt"
	"log/syslog"
	"os"
	"runtime"
	"unsafe"

	"github.com/sirupsen/logrus"
)

func is_64() bool {
	if unsafe.Sizeof(uintptr(0)) == 8 {
		return true
	} else if unsafe.Sizeof(uintptr(0)) == 4 {
		return false
	} else {
		fmt.Println("Unknow system arch")
		os.Exit(-9)
	}
}
func main() {
	logger := logrus.New()
	syslogHook, err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_INFO, "")
	if err != nil {
		logger.Fatal("Unable to create syslog hook:", err)
	}

	fmt.Println("Zygisk run on", runtime.GOARCH)
	zygiskv := os.Getenv("ZKSU_VERSION")
	fmt.Println("zygsik version", zygiskv)
	loadmoudle(runtime.GOARCH)
}
