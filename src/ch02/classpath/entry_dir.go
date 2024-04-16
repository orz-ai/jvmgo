package classpath

import (
	"io/ioutil"
	"path/filepath"
)

/**
 * 表示目录形式的类路径
 */
type DirEntry struct {
	// 存放目录的绝对路径
	absDir string
}

/**
 * @Description: 构造方法
 */
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir}
}

/**
 * @Description: 读取class文件内容
 */
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 先把参数转成绝对路径
	fileName := filepath.Join(self.absDir, className)

	// 读取class文件内容
	file, err := ioutil.ReadFile(fileName)

	return file, self, err
}

func (self *DirEntry) String() string  {
	return self.absDir
}