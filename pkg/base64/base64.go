package base64

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func Process(value string, path string, optionDecode bool, urlEncoding bool) (string, error) {
	var (
		data            = value
		processedString string
		err             error
	)

	if path != "" {
		data = getFileContent(path)
	}

	if optionDecode == false {
		processedString = encode(data, urlEncoding)
	} else {
		processedString, err = decode(data)
		if err != nil {
			return "", err
		}
	}

	return processedString, nil
}

func getFileContent(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}

	return string(data)
}

func decode(value string) (string, error) {
	var decodedString []byte
	var err error

	decodedString, err = b64.StdEncoding.DecodeString(value)

	return string(decodedString), err
}

func encode(value string, urlencoding bool) string {
	var encodedString string

	if urlencoding {
		encodedString = b64.URLEncoding.EncodeToString([]byte(value))
	} else {
		encodedString = b64.StdEncoding.EncodeToString([]byte(value))
	}

	return encodedString
}