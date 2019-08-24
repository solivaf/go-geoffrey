package app

import (
	"errors"
	"github.com/solivaf/go-geoffrey/internal/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	defaultPath  = "./"
	fileFormat   = ".yml"
	baseFileName = "application"
)

const GeoffreyConfig = "GEOFFREY_CONFIG"

func LoadConfig() *config.AppConfig {
	configPath := getConfigPath()
	filePath := getFilePath(configPath)
	fileInfos, _ := ioutil.ReadDir(filePath)
	files := getFilesFromPath(fileInfos)
	b := getFileBytes(filePath, files[0])
	cfg := unmarshal(b)

	return &cfg
}

func getFilesFromPath(fileInfos []os.FileInfo) []string {
	files := make([]string, 0)
	for _, fileInfo := range fileInfos {
		if hasApplicationConfig(fileInfo) {
			files = append(files, fileInfo.Name())
		}
	}
	return files
}

func unmarshal(b []byte) config.AppConfig {
	var c config.AppConfig
	if err := yaml.Unmarshal(b, &c); err != nil {
		panic(errors.New("bad file format. should be yaml valid"))
	}
	return c
}

func getConfigPath() string {
	path := os.Getenv(GeoffreyConfig)
	if len(path) == 0 {
		path = defaultPath
	}
	return path
}

func getFileBytes(path, fileName string) []byte {
	f, err := os.Open(path + "/" + fileName)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return b
}

func getFilePath(path string) string {
	dirLocation := filepath.Dir(path)
	path, err := filepath.Abs(dirLocation)
	if err != nil {
		panic(err)
	}
	return path
}

func hasApplicationConfig(fileInfo os.FileInfo) bool {
	return strings.Contains(fileInfo.Name(), baseFileName) && strings.HasSuffix(fileInfo.Name(), fileFormat)
}
