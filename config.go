package config

import (
	"github.com/ted-escort/utils"
	"io/fs"
	"io/ioutil"
	"os"
)

// Dir the directory to store cache files
func Dir() string {
	return "./data/config/"
}

// FileSuffix cache file suffix. Defaults to '.bin'.
func FileSuffix() string {
	return ".conf"
}

// DirMode the permission to be set for newly created directories.
func DirMode() int {
	return 0775 // 0666
}

// Name the name of cache file
func Name(key string) string {
	return utils.Md5(key)
}

// cacheDirInit 缓存目录初始化
func cacheDirInit() (string, error) {
	// 缓存目录
	fileDir := Dir()
	// 创建目录
	createDir, _ := utils.CreateDir(fileDir)
	if !createDir {
		return "", nil
	}
	return fileDir, nil
}

// File 缓存文件
func File(key string) (string, error) {
	// 缓存目录
	fileDir, err := cacheDirInit()
	if err != nil {
		return "", err
	}
	// 缓存文件
	fileName := Name(key)
	// 后缀
	fileSuffix := FileSuffix()
	// 文件完整路径
	filename := fileDir + fileName + fileSuffix
	if !utils.FileIsExist(filename) {
		_, err := os.Create(filename)
		//_, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fs.FileMode(DirMode()))
		//defer file.Close()
		if err != nil {
			return "", err
		}
	}
	return filename, nil
}

// Get 获取缓存
func Get(key string) ([]byte, error) {
	cacheFile, err := File(key)
	if err != nil {
		return nil, err
	}
	bytes, err := utils.ReadFile(cacheFile)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Set 设置缓存
func Set(key string, value []byte) (bool, error) {
	// 获取缓存文件路径
	cacheFile, err := File(key)
	if err != nil {
		return false, err
	}
	// 写入缓存内容
	err = ioutil.WriteFile(cacheFile, value, fs.FileMode(DirMode()))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Delete 删除缓存
func Delete(key string) error {
	// 获取缓存文件路径
	cacheFile, err := File(key)
	if err != nil {
		return err
	}
	err = os.Remove(cacheFile)
	if err != nil {
		return err
	}
	return nil
}
