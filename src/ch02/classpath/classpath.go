package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

/**
 * @Description: 使用-Xjre选项解析启动类路径和扩展类路径, 使用-classpath/-cp选项来解析用户类路径
 * @param jreOption
 * @param cpOption
 * @return *Classpath
 */
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtraClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

/**
 * @Description: 如果用户没有提供-classpath/-cp选项，则使用当前路径作为类路径
 * @receiver self
 * @param className
 * @return []byte
 * @return Entry
 * @return error
 */
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	return self.userClasspath.readClass(className)
}

/**
 * @Description: 返回用户类路径的字符串表示
 * @receiver self
 * @return string
 */
func (self *Classpath) String() string {
	return self.userClasspath.String()
}

func (self *Classpath) parseBootAndExtraClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// 读取jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// 读取jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	self.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder!")
}

/**
 * @Description: 判断目录是否存在
 * @param path 目录路径
 * @return bool
 */
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
