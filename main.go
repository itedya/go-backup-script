package main

import (
	"log"
	"strings"
	"time"
)

func main() {
	log.SetFlags(0)

	currentTime := time.Now()

	backupFrom, backupTo, duration, err := parseTerminalArguments()

	if err != nil {
		log.Fatal(err)
	}

	archiveTime := strings.Replace(strings.Split(currentTime.Format(time.RFC3339Nano), ".")[0], ":", ".", -1)
	args := []string{"-cvzf", backupTo + archiveTime + "_backup.tar.gz ", backupFrom + "*"}

	_, err = RunCMD("tar", args)

	if err != nil {
		log.Fatal(err)
		return
	}

	if duration != "" {
		err = removeOldFiles(backupTo, duration)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	return
}
