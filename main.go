package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func getCurrentDate() string {
	var date string = fmt.Sprintf("%v%v%v", time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	return date
}

func main() {
	var fileName string = os.Args[1]

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s\n", content)

	fmt.Printf("%v", getCurrentDate())
}
