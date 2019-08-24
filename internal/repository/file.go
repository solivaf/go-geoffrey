package repository

import (
	"io/ioutil"
	"os"
	"strings"
)

const checkoutPath = "/tmp/go-geoffrey/"
const ymlSuffix = ".yml"

var files map[string]string

func init() {
	files = make(map[string]string)
}

func CheckoutPath(name string) string {
	return checkoutPath + name
}

func Files() map[string]string {
	return files
}

func isValidFile(fileName string) bool {
	return strings.HasSuffix(fileName, ymlSuffix) || strings.HasSuffix(fileName, ".yaml")
}

func GetFormattedFileNameWithEnv(appName, appEnv string) string {
	if appEnv == "" {
		return appName + ymlSuffix
	}
	return appName + "-" + appEnv + ymlSuffix
}

func GetFormattedFileName(appName string) string {
	return GetFormattedFileNameWithEnv(appName, "")
}

func GetDefaultFileName() string {
	return "application.yml"
}

func GetDefaultFileNameWithEnv(env string) string {
	return "application-" + env + ymlSuffix
}

func LoadFiles() {
	fileInfos, _ := ioutil.ReadDir(checkoutPath)
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			path := checkoutPath + fileInfo.Name()
			fi, _ := ioutil.ReadDir(path)
			for _, specificFileInfo := range fi {
				if isValidFile(specificFileInfo.Name()) {
					file, _ := os.Open(path + "/" + specificFileInfo.Name())
					bytesContent, _ := ioutil.ReadAll(file)
					files[specificFileInfo.Name()] = string(bytesContent)
				}
			}
		}
	}
}
