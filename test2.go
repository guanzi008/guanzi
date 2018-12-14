package main

import (
	"fmt"
	"github.com/xiaoqidun/goini"
	"log"
)

func main() {
	ini := goini.NewGoINI()
	if err := ini.LoadFile(`H:/go/guanzi008/test2.ini`); err != nil {
		log.Println(err)
		return
	}

	a := ini.GetString("test", "hello", "默认值")
	fmt.Println(a)
}
