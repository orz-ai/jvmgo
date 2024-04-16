package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

/**
 * 表示zip或者jar文件形式的类路径
 */
type ZipEntry struct {
	// 存放文件的绝对路径
	absPath string
}

/**
 * @Description: 相当于ZipEntry的构造函数
 * @param path
 * @return *ZipEntry
 */
func newZipEntry(path string) *ZipEntry {
	// 把路径转成绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}

/**
 * @Description: 从zip文件中读取class文件
 */
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 打开zip文件
	reader, err := zip.OpenReader(self.absPath)
	if err != nil {
		// 打开zip失败直接返回
		return nil, nil, err
	}
	defer reader.Close()

	// 遍历zip里的文件
	for _, f := range reader.File {
		if f.Name == className {
			// 找到指定的class文件，打开
			rc, err := f.Open()
			if err != nil {
				// 打开class文件失败，返回
				return nil, nil, err
			}
			defer rc.Close()

			// 读取所有的
			data, err := ioutil.ReadAll(rc)
			return data, self, nil
		}
	}
	return nil, nil, errors.New("Class not found: " + className)
}

/**
 * @Description: toString方法
 */
func (self *ZipEntry) String() string {
	return self.absPath
}