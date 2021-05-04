package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunCMD(path string, args []string) (out string, err error) {
	fmt.Println("Executing command...")
	cmd := exec.Command(path, args...)

	var b []byte
	b, err = cmd.CombinedOutput()
	out = string(b)

	if err != nil {
		fmt.Println("Executed command \"" + strings.Join(cmd.Args[:], " ") + "\" with error!")
		return "", errors.New(out)
	} else {
		fmt.Println("Executed command \"" + strings.Join(cmd.Args[:], " ") + "\" successfully!")
	}

	return out, nil
}

func parseTerminalArguments() (string, string, string, error) {
	if len(os.Args) != 4 && len(os.Args) != 3 {
		return "", "", "", errors.New("Specify 2 or 3 arguments! Backup From, Backup To and Duration (optional)")
	}

	backupFrom := os.Args[1]
	backupTo := os.Args[2]
	var duration string
	if len(os.Args) == 3 {
		duration = ""
	} else {
		duration = os.Args[3]
	}

	if _, err := os.Stat(backupFrom); os.IsNotExist(err) {
		return "", "", "", errors.New("Backup from path doesn't exist!")
	}

	if _, err := os.Stat(backupTo); os.IsNotExist(err) {
		return "", "", "", errors.New("Backup to path doesn't exist!")
	}

	if backupFrom[len(backupFrom)-1] != 92 && backupFrom[len(backupFrom)-1] != 47 {
		backupFrom += "/"
	}

	if backupTo[len(backupTo)-1] != 92 && backupTo[len(backupTo)-1] != 47 {
		backupTo += "/"
	}

	return backupFrom, backupTo, duration, nil
}
