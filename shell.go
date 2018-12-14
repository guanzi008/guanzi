package main

import (
	"bytes"
	//	"fmt"
	"github.com/xiaoqidun/goini"

	//  "fmt"
	"log"
	"os/exec"
)

func exec_shell(s string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	//    fmt.Printf("%s", out.String())
}

func main() {
	ini := goini.NewGoINI()
	if err := ini.LoadFile(`H:/go/guanzi008/test2.ini`); err != nil {
		log.Println(err)
		return
	}

	a := ini.GetString("test", "hello", "默认值")
	exec_shell("a" + a)
}
