package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func removeOldFiles(backupTo string, duration string) error {
	files, err := filepath.Glob(backupTo + "*")
	if err != nil {
		return errors.New("can't read files in backup to filepath")
	}

	filesToCheck := map[int]string{}

	durationParsed, err := time.ParseDuration(duration)
	if err != nil {
		return errors.New("can't parse date")
	}

	timeToRemove, err := time.Parse("2006-01-02-15-04-05", time.Now().Add(durationParsed).Format("2006-01-02-15-04-05"))
	if err != nil {
		return errors.New("can't parse date")
	}

	for _, file := range files {
		fileBase := filepath.Base(file)
		fileBase = strings.Replace(fileBase, "_backup.tar.gz", "", -1)

		fileTime, err := time.Parse("2006-01-02-15-04-05", fileBase)
		if err == nil {
			filesToCheck[len(filesToCheck)] = file

			if fileTime.Unix() <= timeToRemove.Unix() {
				err := os.Remove(file)
				if err != nil {
					return err
				}
				fmt.Println("Deleted " + file)
			}
		}
	}

	return nil
}
