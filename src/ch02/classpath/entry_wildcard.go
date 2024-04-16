package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	// 去除路径末尾的*
	baseDir := path[:len(path)-1]
	compositeEntry := [] Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			// 跳过子目录，通配符不能递归匹配子目录以下的文件
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}

		return nil
	}


	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
