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
	ini.GetNames("")
	a := ini.GetString("test", "url", "默认值")

	ini.LoadFile("H:/go/guanzi008/test2.ini")
	fmt.Println(ini.GetNames(""))
	fmt.Println(ini.GetString("a", "a", "a"))
	fmt.Println(a)

}
func main1() {
	ini := goini.NewGoINI()
	if err := ini.LoadFile("./test.ini"); err != nil {
		log.Println(err)
		return
	}
	fmt.Println(ini.GetString("a", "b", "默认值"))
}
