package utils

import (
	"encoding/json"
	"os"
)

func Restore(fileName string, BackupVersionFile string) {
	// Get current backup file status from `file_version.json`
	fileVersions := map[string][]string{}

	versionData, err := os.ReadFile(BackupVersionFile)
	checkIfFatal(err)

	err = json.Unmarshal(versionData, &fileVersions)
	checkIfFatal(err)

	backupFileName := fileVersions[fileName][len(fileVersions[fileName])-1]

	// Exchange data between the current file and the latest backup file
	fileData, err := os.ReadFile(fileName)
	checkIfFatal(err)

	backupFileData, err := os.ReadFile(backupFileName)
	checkIfFatal(err)

	err = os.WriteFile(fileName, backupFileData, FilePermission)
	checkIfFatal(err)

	err = os.WriteFile(backupFileName, fileData, FilePermission)
	checkIfFatal(err)
}
