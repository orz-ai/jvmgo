package main

import (
	"ch03/classpath"
	"fmt"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("Version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	// 解析出命令行参数
	cp := classpath.Parse(cmd.XJreOption, cmd.cpOption)
	fmt.Printf(
		"classpath:%s class:%s args:%v\n",
		cmd.cpOption,
		cmd.class,
		cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err :=cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data : %v\n", classData)
}
