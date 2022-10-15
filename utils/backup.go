package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const FilePermission = 0600

func checkIfFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isNotContained(versions []string, oldVersion string) bool {
	for _, version := range versions {
		if version == oldVersion {
			return false
		}
	}
	return true
}

func getCurrentDate() string {
	var date string = fmt.Sprintf("%v%v%v", time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	return date
}

func Backup(fileName string, BackupVersionFile string) {
	backupFileName := fileName + "." + getCurrentDate() + ".backup"

	// Backup file
	fileData, err := ioutil.ReadFile(fileName)
	checkIfFatal(err)

	err = os.WriteFile(backupFileName, fileData, FilePermission)
	checkIfFatal(err)

	// Get current backup file status from `file_version.json`
	fileVersions := map[string][]string{}

	versionData, err := os.ReadFile(BackupVersionFile)
	checkIfFatal(err)

	err = json.Unmarshal(versionData, &fileVersions)
	checkIfFatal(err)

	// Check if the new backup file name is existing in the list of backup files or not
	// and update `file_version.json`
	if isNotContained(fileVersions[fileName], backupFileName) {
		fileVersions[fileName] = append(fileVersions[fileName], backupFileName)
	}

	newVersion, err := json.Marshal(fileVersions)
	checkIfFatal(err)

	err = os.WriteFile(BackupVersionFile, newVersion, FilePermission)
	checkIfFatal(err)
}
