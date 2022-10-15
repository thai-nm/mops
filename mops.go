package main

import (
	"fmt"
	"os"

	"github.com/thainmuet/backup/utils"
)

const BackupVersionFile = "backup_versions.json"

func main() {

	subcommand := os.Args[1]
	fileName := os.Args[2]

	switch subcommand {
	case "backup":
		utils.Backup(fileName, BackupVersionFile)
	case "restore":
		utils.Restore(fileName, BackupVersionFile)
	default:
		fmt.Printf("There is no '%s' subcommand.\n", subcommand)
	}
}
