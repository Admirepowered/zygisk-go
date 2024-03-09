package main

import (
	"log"

	"github.com/Admirepowered/zygisk-go/util"
)

func Loadmoudle(arch string) {
	return
}
func Start() {
	log.Println("Running")
	util.Switch_mount_namespace(1)

}
