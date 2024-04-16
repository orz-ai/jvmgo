package classpath

import (
	"os"
	"strings"
)

// 存放文件分隔符 ";"
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {

	// 负责寻找和加载class文件
	readClass(className string) ([]byte, Entry, error)

	// 相当于java中的toString
	String() string

}

/**
 * @Description: 创建不同的entry实例
 * @param path
 * @return Entry
 */
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*"){
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP"){
		return newZipEntry(path)
	}

	return newDirEntry(path)
}