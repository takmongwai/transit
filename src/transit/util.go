package main

import (
	"os"
	"path/filepath"
)

//当前运行目录
func SelfRoot() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logger.Fatal(err)
	}
	return dir
}

//判断文件是否存在
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//判断参数是否为空
func IsEmpty(v interface{}) bool {
	return false
}
