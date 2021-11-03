package util

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// GetDefaultFilePath 获得默认地址
func GetDefaultFilePath() string {
	dir, _ := os.Getwd()
	path := dir + fmt.Sprintf("/tmp/")
	return path
}

// GetDefaultFileName 获得默认文件名
func GetDefaultFileName() string {
	return fmt.Sprintf("%14v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000000000))
}

// CreateTmpDir 创建文件夹
func CreateTmpDir(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		//创建文件夹
		err = os.Mkdir(path, os.ModePerm)
		return err
	}
	return err
}