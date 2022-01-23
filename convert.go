package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type ErrorLog struct {
	DateTime   string
	Status     string
	UniqueCode string
	Reason     string
}

func HandleConvertToJson(locFile, locWriteFile string) {
	b, err := ioutil.ReadFile(locFile) // just pass the file name
	if err != nil {
		panic(err.Error())
	}

	locWriteFile = HandleReplaceFileType(locWriteFile, "json")

	var dataLog = make([]ErrorLog, 0)
	str := string(b) // convert content to a 'string'
	lines := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")

	for _, line := range lines {
		val := strings.Split(line, " ")
		reason := strings.Join(val[4:], " ")
		dataLog = append(dataLog, ErrorLog{
			DateTime:   fmt.Sprintf("%s %s", val[0], val[1]),
			Status:     val[2],
			UniqueCode: val[3],
			Reason:     reason,
		})
	}

	file, err := json.MarshalIndent(dataLog, "", " ")

	if err != nil {
		panic(err.Error())
	}

	err = ioutil.WriteFile(locWriteFile, file, 0644)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success convert file to JSON", locWriteFile)

}

func HandleConvertToText(locReadFile, locWriteFile string) {
	b, err := ioutil.ReadFile(locReadFile) // just pass the file name

	locWriteFile = HandleReplaceFileType(locWriteFile, "text")
	if err != nil {
		panic(err.Error())
	}

	err = ioutil.WriteFile(locWriteFile, b, 0644)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("success convert file to TEXT", locWriteFile)
}

// for handle name and dot not same
func HandleReplaceFileType(locWriteFile, typeFile string) string {
	splitLocWrite := strings.Split(locWriteFile, "/")
	newSplitDot := strings.Split(splitLocWrite[len(splitLocWrite)-1], ".")
	if len(newSplitDot) > 1 {
		if newSplitDot[len(newSplitDot)-1] != typeFile {
			newSplitDot[len(newSplitDot)-1] = typeFile
			splitLocWrite[len(splitLocWrite)-1] = strings.Join(newSplitDot, ".")
			locWriteFile = strings.Join(splitLocWrite, "/")
		}
	} else {
		panic("cannot convert file")
	}

	return locWriteFile
}
