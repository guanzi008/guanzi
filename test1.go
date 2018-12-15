package main

import (
	"fmt"
	"github.com/xiaoqidun/goini"
	//	"net/url"
)

func main() {
	ini := goini.NewGoINI()
	if err := ini.LoadFile(`H:\go\guanzi008\test2.ini`); err != nil {
		fmt.Println("配置解析错误", err)
		return
	} else {
		fmt.Println(ini.GetNames(""))
		fmt.Println(ini.GetString("hi", "url", ""))
		fmt.Println(ini.GetNames(""))
		fmt.Println(ini.GetString("a", "a", "a"))
	}
}

//GetString
//第一个参数是[xxx]，中括号中的值，第二个参数是[xxx]这个分区下的键名称，比如hello = world，那么第二个参数就是hello，第三个是在这个值配置文件中没有找到时使用的默认值
