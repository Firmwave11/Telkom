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

type AccessLog struct {
	Ip       string
	DateTime string
	Reason   string
}

func HandleConvertToJson(locFile, locWriteFile string) {
	b, err := ioutil.ReadFile(locFile) // just pass the file name
	if err != nil {
		panic(err.Error())
	}

	locWriteFile = HandleReplaceFileType(locWriteFile, "json")
	logType := GetFileType(locFile)

	var file []byte
	str := string(b) // convert content to a 'string'
	lines := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")

	switch logType {
	case "error":
		file = HandleErrorLog(lines)
	case "access":
		file = HandleAccessLog(lines)
	default:
		panic("cannot handle this file !! this program just handle convert json for access.log or error.log")
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

func HandleErrorLog(lines []string) []byte {
	var errorLog = make([]ErrorLog, 0)
	for _, line := range lines {
		val := strings.Split(line, " ")
		if len(val) > 4 {
			reason := strings.Join(val[4:], " ")
			errorLog = append(errorLog, ErrorLog{
				DateTime:   fmt.Sprintf("%s %s", val[0], val[1]),
				Status:     val[2],
				UniqueCode: val[3],
				Reason:     reason,
			})
		}
	}
	file, err := json.MarshalIndent(errorLog, "", " ")

	if err != nil {
		panic(err.Error())
	}

	return file
}

func HandleAccessLog(lines []string) []byte {
	var accessLog = make([]AccessLog, 0)
	for _, line := range lines {
		val := strings.Split(line, " ")
		if len(val) > 5 {
			reason := strings.Join(val[5:], " ")
			accessLog = append(accessLog, AccessLog{
				Ip:       val[0],
				DateTime: strings.ReplaceAll(val[3], "[", ""),
				Reason:   reason,
			})
		}

	}
	file, err := json.MarshalIndent(accessLog, "", " ")

	if err != nil {
		panic(err.Error())
	}

	return file
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

func GetFileType(locReadFile string) string {
	splitLocWrite := strings.Split(locReadFile, "/")
	newSplitDot := strings.Split(splitLocWrite[len(splitLocWrite)-1], ".")
	logType := newSplitDot[0]

	return logType
}
