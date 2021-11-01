package service

import (
	"douyin/api"
	"strings"
)

func ApiHandler(url string) (r string) {

	if strings.Contains(url, "douyin") { // 没有匹配时，值为-1
		r = api.DouYin(url)
	} else if strings.Contains(url, "weishi") {
		r = api.WeiShi(url)
	} else if strings.Contains(url, "pipix") {
		r = api.PiPix(url)
	} else if strings.Contains(url, "kuaishou") {
		r = api.KuaiShou(url)
	} else if strings.Contains(url, "kg3.qq.com") {
		r = api.Kg3(url)
	} else if strings.Contains(url, "ixigua.com") {
		r = "截至2021-04-04，仅能通过下载合并的方式获取无无水印视频"
	} else if strings.Contains(url, "eyepetizer") {
		r = api.EyePetizer(url)
	} else if strings.Contains(url, "vuevideo") {
		r = api.VueVlog(url)
	} else if strings.Contains(url, "xiaokaxiu") {
		r = "暂未支持该接口，请提交issue"
	} else if strings.Contains(url, "ippzone") {
		r = api.IppZone(url)
	} else if strings.Contains(url, "weibo.com") {
		r = api.WeiBo(url)
	} else if strings.Contains(url, "zuiyou") {
		r = api.ZuiYou(url)
	} else if strings.Contains(url, "bbq.bilibili") {
		r = api.QingShi(url)
	} else if strings.Contains(url, "bilibili.com") {
		r = api.Bilibili(url)
	} else if strings.Contains(url, "immomo") {
		r = api.MonMo(url)
	} else {
		r = "暂未支持该接口，请提交issue"
	}

	return 
}

