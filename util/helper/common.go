package helper

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/kaifei-bianjie/mock/conf"
	"github.com/kaifei-bianjie/mock/util/constants"
	"github.com/parnurzeal/gorequest"
	"os"
	"strconv"
	"strings"
)

// post json data use http client
func HttpClientPostJsonData(uri string, requestBody *bytes.Buffer) (int, []byte, error) {
	var (
		errStrs    []string
		err        error
		statusCode int
	)
	url := conf.NodeUrl + uri
	reqStr := string(requestBody.Bytes())

	resp, body, errs := gorequest.New().Timeout(constants.HttpTimeout).Post(url).
		Send(reqStr).EndBytes()

	if len(errs) > 0 {
		for _, err := range errs {
			errStrs = append(errStrs, err.Error())
		}
		err = fmt.Errorf(strings.Join(errStrs, "|"))
	} else {
		statusCode = resp.StatusCode
	}

	return statusCode, body, err
}

// get data use http client
func HttpClientGetData(uri string) (int, []byte, error) {
	var (
		err        error
		errStrs    []string
		statusCode int
	)
	url := conf.NodeUrl + uri

	resp, body, errs := gorequest.New().Timeout(constants.HttpTimeout).Get(url).
		//Retry(2, 5*time.Second, http.StatusInternalServerError).
		EndBytes()

	if len(errs) > 0 {
		for _, err := range errs {
			errStrs = append(errStrs, err.Error())
		}
		err = fmt.Errorf(strings.Join(errStrs, "|"))
	} else {
		statusCode = resp.StatusCode
	}

	return statusCode, body, err
}

func ConvertStrToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// check file whether exist
// return true if exist, otherwise return false
func CheckFileExist(filePath string) (bool, error) {
	exists := true
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			exists = false
		} else {
			// unknown err
			return false, err
		}
	}
	return exists, nil
}

// create folder if not exist
// return err if not successful create
func CreateFolder(folderPath string) error {
	folderExist, err := CheckFileExist(folderPath)
	if err != nil {
		return err
	}

	if !folderExist {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteFile(filePath string, content []byte) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	fileWrite := bufio.NewWriter(file)
	_, err = fileWrite.Write(content)
	if err != nil {
		return err
	}
	fileWrite.Flush()
	return nil
}
