package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func RunCMD(path string, args []string, debug bool) (out string, err error) {

	cmd := exec.Command(path, args...)

	var b []byte
	b, err = cmd.CombinedOutput()
	out = string(b)

	if debug {
		fmt.Println(strings.Join(cmd.Args[:], " "))

		if err != nil {
			fmt.Println("RunCMD ERROR")
			fmt.Println(out)
		}
	}

	return
}

func Tar(source string, backupTo string) error {
	currentTime := time.Now()

	filename := backupTo + currentTime.Format("2006-01-02-15-04-05") + "_backup.tar.gz"

	lastIndex := strings.LastIndex(source, string(filepath.Separator))
	args := []string{"-czf", filename, "-C", source[:lastIndex], source[lastIndex+1:]}

	_, err := RunCMD("tar", args, true)

	if err != nil {
		return err
	}

	return nil
}
