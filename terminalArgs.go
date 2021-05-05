package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func filterPath(path string) (string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errors.New("backup from path doesn't exist")
	}

	path = strings.ReplaceAll(path, "/", string(filepath.Separator))
	path = strings.ReplaceAll(path, "\\", string(filepath.Separator))

	if !strings.HasSuffix(path, string(filepath.Separator)) {
		path += string(filepath.Separator)
	}

	return path, nil
}

func parseTerminalArguments() (string, string, string, error) {
	if len(os.Args) != 4 && len(os.Args) != 3 {
		return "", "", "", errors.New("specify 2 or 3 arguments: backup from, backup to and duration (optional)")
	}

	backupFrom, err := filterPath(os.Args[1])
	if err != nil {
		return "", "", "", err
	}
	backupFrom = backupFrom[:len(backupFrom)-1]

	backupTo, err := filterPath(os.Args[2])
	if err != nil {
		return "", "", "", err
	}

	var duration string
	if len(os.Args) == 3 {
		duration = ""
	} else {
		duration = os.Args[3]
		_, err = time.ParseDuration(duration)
		if err != nil {
			return "", "", "", errors.New("can't parse duration")
		}
	}

	return backupFrom, backupTo, duration, nil
}
