package service

import (
	"binginx.com/brush/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func assembleHeader(req *http.Request) {
	req.Header.Add("Host", "ecosys-web.china-inv.cn")
	req.Header.Add("Accept-Encoding", "gzip,deflate")
	req.Header.Add("sign", "fbbee325b38811c5bf83c06df62851e2")
	req.Header.Add("Referer", "https://ecosys-web.china-inv.cn/index.html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Set("Content-Type", "application/json")
}

func ScoreTotal(user model.UserInfo) (err error, score string) {
	log.Println("获取分数==========================")
	req, err := http.NewRequest("GET", "https://ecosys-web.china-inv.cn/api/score/scoreTotal?qtime="+strconv.Itoa(time.Now().Nanosecond()), nil)
	if err != nil {
		log.Println(err)
		return
	}
	assembleHeader(req)
	req.Header.Set("accessToken", user.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	var scoreTotal model.Resp
	err = json.Unmarshal(out, &scoreTotal)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(scoreTotal)
	return err, scoreTotal.Data
}

func EcosysNews(user model.UserInfo) (err error, ecosysNewsId []model.Content) {

	req, err := http.NewRequest("GET", "https://ecosys-web.china-inv.cn/api/content/list?qtime="+strconv.Itoa(time.Now().Nanosecond())+"&columnId=unEcosysNews&pageNum=1&pageSize=500&needAll=true", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	assembleHeader(req)
	req.Header.Set("accessToken", user.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	log.Println("response Status:", resp.Status)
	var ecosysNews model.Contents
	err = json.Unmarshal(out, &ecosysNews)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(ecosysNews)
	return err, ecosysNews.Data
}

func ViewNews(user model.UserInfo, news []model.Content) (err error) {
	for _, value := range news {
		req, err1 := http.NewRequest("GET", "https://ecosys-web.china-inv.cn/api/content?qtime="+strconv.Itoa(time.Now().Nanosecond())+"&contentId="+strconv.Itoa(value.ContentId), nil)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		assembleHeader(req)
		req.Header.Set("accessToken", user.Token)
		client := &http.Client{}
		resp, err2 := client.Do(req)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		defer resp.Body.Close()
		//out, err3 := ioutil.ReadAll(resp.Body)
		//if err3 != nil {
		//	fmt.Println(err3)
		//	return
		//}
		log.Println("response Status:", resp.Status)

	}
	return
}

func Behavior(user model.UserInfo, flag string, news []model.Content) (err error) {
	for _, value := range news {
		log.Println(value.ContentId)
		behavior := &model.Behavior{
			UserBehavior: flag,
			ResourceId:   value.ContentId,
		}
		data, err1 := json.Marshal(behavior)
		if err1 != nil {
			return
		}
		log.Println(string(data))
		form := url.Values{}
		form.Add("userBehavior", flag)
		form.Add("resourceId", strconv.Itoa(value.ContentId))
		req, err2 := http.NewRequest("POST", "https://ecosys-web.china-inv.cn/api/content/independent/content/behavior?userBehavior="+flag+"&resourceId="+strconv.Itoa(value.ContentId), nil)
		if err2 != nil {
			return
		}
		assembleHeader(req)
		req.Header.Add("accessToken", user.Token)
		client := &http.Client{}
		resp, err3 := client.Do(req)
		if err3 != nil {
			return
		}
		defer resp.Body.Close()
		//out, err4 := ioutil.ReadAll(resp.Body)
		//if err4 != nil {
		//	return
		//}
		log.Println("response Status:", resp.Status)
		duration := time.Duration(2) * time.Second
		time.Sleep(duration)
	}
	return
}
