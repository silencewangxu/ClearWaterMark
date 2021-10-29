package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var url string

func main() {
	flag.StringVar(&url, "url", "", "分享链接")
	flag.Parse()
	if url == "" || find(url) == "" {
		fmt.Println("无效的地址")
		os.Exit(0)
	}
	fmt.Printf("去水印后抖音地址为:%s", getVideoUrl(url))
}
func find(url string) string {
	reg := regexp.MustCompile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
	result := reg.FindAllStringSubmatch(url, 1)
	if len(result) == 0 {
		return ""
	}
	return result[0][0]
}

func getVideoUrl(url string) string {
	m := make(map[string]string)
	response := httpGet(url, m)
	response.Body.Close()
	var respStr = response.Request.URL.String()
	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(`video/(\d+)?`)
	if reg == nil { //解释失败，返回nil
		fmt.Println("regexp err")
		return ""
	}
	//根据规则提取关键信息
	result := reg.FindAllStringSubmatch(respStr, 1)
	//todo 这里要加判断
	itemsId := result[0][1]
	m["Content-Type"] = "application/json"
	var f List
	resp := httpGet("https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids="+itemsId, m)
	body, _ := ioutil.ReadAll(resp.Body)

	response.Body.Close()
	json.Unmarshal(body, &f)
	s := f.ItemList[0].Video.PlayAddr.URLList[0]
	return strings.Replace(s, "playwm", "play", 1)

}

func httpGet(url string, headers map[string]string) *http.Response {
	client := &http.Client{}
	//提交请求
	request, err := http.NewRequest("GET", url, nil)
	//增加header选项
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Mobile Safari/537.36 Edg/87.0.664.66")
	if len(headers) != 0 {
		for key, header := range headers {
			request.Header.Add(key, header)
		}
	}
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)

	return response
}

type List struct {
	ItemList []struct {
		Music struct {
			Duration   int    `json:"duration"`
			ID         int64  `json:"id"`
			Mid        string `json:"mid"`
			Title      string `json:"title"`
			Author     string `json:"author"`
			CoverLarge struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_large"`
			CoverMedium struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_medium"`
			PlayURL struct {
				URLList []string `json:"url_list"`
				URI     string   `json:"uri"`
			} `json:"play_url"`
			Position interface{} `json:"position"`
			Status   int         `json:"status"`
			CoverHd  struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_hd"`
			CoverThumb struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_thumb"`
		} `json:"music"`
		RiskInfos struct {
			Type    int    `json:"type"`
			Content string `json:"content"`
			Warn    bool   `json:"warn"`
		} `json:"risk_infos"`
		VideoText  interface{} `json:"video_text"`
		Promotions interface{} `json:"promotions"`
		IsPreview  int         `json:"is_preview"`
		CreateTime int         `json:"create_time"`
		ShareURL   string      `json:"share_url"`
		Statistics struct {
			AwemeID      string `json:"aweme_id"`
			CommentCount int    `json:"comment_count"`
			DiggCount    int    `json:"digg_count"`
			PlayCount    int    `json:"play_count"`
			ShareCount   int    `json:"share_count"`
		} `json:"statistics"`
		LabelTopText interface{} `json:"label_top_text"`
		LongVideo    interface{} `json:"long_video"`
		TextExtra    []struct {
			HashtagName string `json:"hashtag_name"`
			HashtagID   int64  `json:"hashtag_id"`
			Start       int    `json:"start"`
			End         int    `json:"end"`
			Type        int    `json:"type"`
		} `json:"text_extra"`
		Geofencing   interface{} `json:"geofencing"`
		IsLiveReplay bool        `json:"is_live_replay"`
		ForwardID    string      `json:"forward_id"`
		ChaList      []struct {
			ChaName        string      `json:"cha_name"`
			UserCount      int         `json:"user_count"`
			ConnectMusic   interface{} `json:"connect_music"`
			HashTagProfile string      `json:"hash_tag_profile"`
			Cid            string      `json:"cid"`
			Type           int         `json:"type"`
			ViewCount      int         `json:"view_count"`
			IsCommerce     bool        `json:"is_commerce"`
			Desc           string      `json:"desc"`
		} `json:"cha_list"`
		AwemeType   int         `json:"aweme_type"`
		GroupID     int64       `json:"group_id"`
		Duration    int         `json:"duration"`
		ImageInfos  interface{} `json:"image_infos"`
		VideoLabels interface{} `json:"video_labels"`
		AwemeID     string      `json:"aweme_id"`
		Desc        string      `json:"desc"`
		Video       struct {
			OriginCover struct {
				URLList []string `json:"url_list"`
				URI     string   `json:"uri"`
			} `json:"origin_cover"`
			Ratio    string `json:"ratio"`
			Vid      string `json:"vid"`
			PlayAddr struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"play_addr"`
			Height   int         `json:"height"`
			Width    int         `json:"width"`
			BitRate  interface{} `json:"bit_rate"`
			Duration int         `json:"duration"`
			Cover    struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover"`
			DynamicCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"dynamic_cover"`
			HasWatermark bool `json:"has_watermark"`
		} `json:"video"`
		CommentList interface{} `json:"comment_list"`
		Images      interface{} `json:"images"`
		Author      struct {
			UID          string `json:"uid"`
			Nickname     string `json:"nickname"`
			AvatarLarger struct {
				URLList []string `json:"url_list"`
				URI     string   `json:"uri"`
			} `json:"avatar_larger"`
			FollowersDetail interface{} `json:"followers_detail"`
			PolicyVersion   interface{} `json:"policy_version"`
			TypeLabel       interface{} `json:"type_label"`
			ShortID         string      `json:"short_id"`
			Signature       string      `json:"signature"`
			AvatarThumb     struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_thumb"`
			AvatarMedium struct {
				URLList []string `json:"url_list"`
				URI     string   `json:"uri"`
			} `json:"avatar_medium"`
			UniqueID         string      `json:"unique_id"`
			PlatformSyncInfo interface{} `json:"platform_sync_info"`
			Geofencing       interface{} `json:"geofencing"`
		} `json:"author"`
		ShareInfo struct {
			ShareDesc      string `json:"share_desc"`
			ShareTitle     string `json:"share_title"`
			ShareWeiboDesc string `json:"share_weibo_desc"`
		} `json:"share_info"`
		AuthorUserID int64 `json:"author_user_id"`
	} `json:"item_list"`
	Extra struct {
		Logid string `json:"logid"`
		Now   int64  `json:"now"`
	} `json:"extra"`
	StatusCode int `json:"status_code"`
}
