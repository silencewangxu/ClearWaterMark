package main

import (
	"douyin/service"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var url string

func main() {
	flag.StringVar(&url, "url", "", "分享链接")
	flag.Parse()

	if url == "" || find(url) == "" {
		fmt.Println("无效的地址")
		os.Exit(0)
	}
	fmt.Println(service.ApiHandler(url))
}
func find(url string) string {
	reg := regexp.MustCompile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
	result := reg.FindAllStringSubmatch(url, 1)
	if len(result) == 0 {
		return ""
	}
	return result[0][0]
}


