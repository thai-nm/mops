package utils

import (
	"encoding/json"
	"os"
)

const FilePermission = 0600

func isNotContained(versions []string, oldVersion string) bool {
	for _, version := range versions {
		if version == oldVersion {
			return false
		}
	}
	return true
}

func Backup(fileName string, BackupVersionFile string) {
	backupFileName := fileName + "." + getCurrentDate() + ".backup"

	// Backup file
	fileData, err := os.ReadFile(fileName)
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
