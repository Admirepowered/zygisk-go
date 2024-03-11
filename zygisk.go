package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Admirepowered/zygisk-go/util"
)

func getSubDirectories(directory string) ([]string, error) {
	var subDirs []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() && path != directory {
			subDirs = append(subDirs, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return subDirs, nil
}
func create_library_fd(sofile string) int {
	return 0
}
func Loadmoudle(arch string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		return
	}

	// 获取上级目录
	parentDir := filepath.Dir(currentDir)
	log.Println("Parent directory:", parentDir)

	subDirs, err := getSubDirectories(parentDir)
	if err != nil {
		log.Println("Error getting subdirectories:", err)
		return
	}
	log.Println("Subdirectories of", parentDir, ":")
	for _, dir := range subDirs {

		log.Println("loaging", dir)
		filePath := dir + "/zygisk/" + Get_arch() + ".so"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Fatal("loading failed")
		} else {
			create_library_fd(filePath)
		}

	}

	return
}
func Start() {
	log.Println("Running")
	util.Switch_mount_namespace(1)

}
