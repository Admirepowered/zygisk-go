package util

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func Switch_mount_namespace(pid int) error {

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	mnt, err := os.Open(fmt.Sprintf("/proc/%d/ns/mnt", pid))
	if err != nil {
		return err
	}

	if err := unix.Setns(int(mnt.Fd()), unix.CLONE_NEWNS); err != nil {
		return err
	}

	if err := os.Chdir(cwd); err != nil {
		return err
	}

	return nil
}
