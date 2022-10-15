package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const VersionFile = "file_versions.json"

type Version struct {
	FileName string
	Versions []string
}

func crash(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getCurrentDate() string {
	var date string = fmt.Sprintf("%v%v%v", time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	return date
}

func getVersionsFromFile(versionFile string) []byte {
	versions, err := ioutil.ReadFile(versionFile)
	crash(err)
	return versions
}

func isNotContained(versions []string, oldVersion string) bool {
	for _, version := range versions {
		if version == oldVersion {
			return false
		}
	}
	return true
}

func main() {
	var fileName string = os.Args[1]
	var backupFileName string = fileName + "." + getCurrentDate() + ".backup"

	fileVersions := []Version{}

	fileContent, err := ioutil.ReadFile(fileName)
	crash(err)

	err = os.WriteFile(backupFileName, fileContent, 0600)
	crash(err)

	err = json.Unmarshal(getVersionsFromFile(VersionFile), &fileVersions)
	crash(err)

	fmt.Printf("%v\n", fileVersions)

	for i := 0; i < len(fileVersions); i++ {
		if fileVersions[i].FileName == fileName && isNotContained(fileVersions[i].Versions, backupFileName) {
			fileVersions[i].Versions = append(fileVersions[i].Versions, backupFileName)
			break
		}
	}

	fmt.Printf("%v\n", fileVersions)

	newVersion, err := json.Marshal(fileVersions)
	crash(err)

	err = os.WriteFile(VersionFile, newVersion, 0600)
	crash(err)
}
