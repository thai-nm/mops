package utils

import (
	"fmt"
	"log"
	"time"
)

func checkIfFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func getCurrentDate() string {
	var date string = fmt.Sprintf("%v%v%v", time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	return date
}
