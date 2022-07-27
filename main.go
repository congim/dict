package main

import (
	"fmt"
	"regexp"
)

func sss() {
	// test := "a,b,c,d,e"
	aaa := "我是测试代码AAA"
	keyWords := []string{"AAA", "BBB", "ccc"}
	for _, v := range keyWords {
		reg := regexp.MustCompile("(?i)" + v)
		fmt.Println(reg)
		aaa = reg.ReplaceAllString(aaa, "*")
		fmt.Println(aaa)
	}
}
func main() {
	// dic := dict.New()
	// dic.AddWorld("DDDD", "fuck")
	// fmt.Println(dic.Replace("ddddddddd Fuck", '*'))
	sss()
}
