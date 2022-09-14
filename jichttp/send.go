package jichttp

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

func ScoreTotal(user model.UserInfo) (err error, score string) {
	log.Println("获取分数==========================")
	req, err := http.NewRequest("GET", "https://ecosys-web.china-inv.cn/api/score/scoreTotal?qtime="+string(time.Now().UnixNano()/1e6), nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Host", "ecosys-web.china-inv.cn")
	req.Header.Add("Accept-Encoding", "gzip,deflate")
	req.Header.Add("sign", "fbbee325b38811c5bf83c06df62851e2")
	req.Header.Add("Referer", "https://ecosys-web.china-inv.cn/index.html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Set("accessToken", user.Token)
	// 再设置个json
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(out))
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
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

	req, err := http.NewRequest("GET", "https://ecosys-web.china-inv.cn/api/content/list?qtime="+string(time.Now().UnixNano()/1e6)+"&columnId=unEcosysNews&pageNum=1&pageSize=200&needAll=true", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Host", "ecosys-web.china-inv.cn")
	req.Header.Add("Accept-Encoding", "gzip,deflate")
	req.Header.Add("sign", "fbbee325b38811c5bf83c06df62851e2")
	req.Header.Add("Referer", "https://ecosys-web.china-inv.cn/index.html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Set("accessToken", user.Token)
	// 再设置个json
	req.Header.Set("Content-Type", "application/json")
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
		req, err1 := http.NewRequest("GET", "https://ecosys-web.china-inv.cn/api/content?qtime="+string(time.Now().UnixNano()/1e6)+"&contentId="+strconv.Itoa(value.ContentId), nil)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		req.Header.Add("Host", "ecosys-web.china-inv.cn")
		req.Header.Add("Accept-Encoding", "gzip,deflate")
		req.Header.Add("sign", "fbbee325b38811c5bf83c06df62851e2")
		req.Header.Add("Referer", "https://ecosys-web.china-inv.cn/index.html")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
		req.Header.Add("Connection", "Keep-Alive")
		req.Header.Set("accessToken", user.Token)
		// 再设置个json
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err2 := client.Do(req)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		defer resp.Body.Close()
		out, err3 := ioutil.ReadAll(resp.Body)
		if err3 != nil {
			fmt.Println(err3)
			return
		}
		log.Println("response Status:", resp.Status)
		log.Println("response Body:", string(out))

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
		//var r http.Request
		//r.ParseForm()
		//r.Form.Add("userBehavior", flag)
		//r.Form.Add("resourceId", strconv.Itoa(value.ContentId))
		//bodystr := strings.TrimSpace(r.Form.Encode())
		form := url.Values{}
		form.Add("userBehavior", flag)
		form.Add("resourceId", strconv.Itoa(value.ContentId))
		req, err2 := http.NewRequest("POST", "https://ecosys-web.china-inv.cn/api/content/independent/content/behavior?userBehavior="+flag+"&resourceId="+strconv.Itoa(value.ContentId), nil)
		//req, err2 := http.NewRequest("POST", "https://ecosys-web.china-inv.cn/api/content/independent/content/behavior", nil)
		if err2 != nil {
			return
		}
		//req.PostForm = form
		req.Header.Add("Host", "ecosys-web.china-inv.cn")
		req.Header.Add("Accept-Encoding", "gzip,deflate")
		req.Header.Add("sign", "d41d8cd98f00b204e9800998ecf8427e")
		req.Header.Add("Referer", "https://ecosys-web.china-inv.cn/index.html")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
		req.Header.Add("Connection", "Keep-Alive")

		//req.Header.Set("accessToken", "eyJhbGciOiJSUzI1NiIsInR5cCI6ImF0K2p3dCJ9.eyJuYmYiOjE2NjE5OTU2MjcsImV4cCI6MTY2MTk5OTIyNywiaXNzIjoiaHR0cHM6Ly9pZHMuY2hpbmEtaW52LmNuIiwiYXVkIjpbImVjb3N5cyIsImVjb3N5c3JjIiwic2VuZGNvZGUiLCJzdmMiXSwiY2xpZW50X2lkIjoiZWNvc3lzIiwic3ViIjoibGliaW5Ac2RqaWN0ZWMuY29tIiwiYXV0aF90aW1lIjoxNjYxOTk1NjI3LCJpZHAiOiJsb2NhbCIsIm5hbWUiOiLmnY7lvawiLCJlbWFpbCI6ImxpYmluQHNkamljdGVjLmNvbSIsImRlcGFydG1lbnQiOiLkuLTml7bpg6jpl6giLCJjb21wYW55Ijoi5bu65oqV5pWw5o2u56eR5oqA77yI5bGx5Lic77yJ5pyJ6ZmQ5YWs5Y-4IiwiZGl2aXNpb24iOiIiLCJlbXBsb3llZWlkIjoiMTIzODk1IiwiZW1haWxfdmVyaWZpZWQiOiJ0cnVlIiwianRpIjoiQkRCRDBGNkI4QTQxNDk5NUU4NTUxNjY4RTM0NkRBN0IiLCJpYXQiOjE2NjE5OTU2MjcsInNjb3BlIjpbImVjb3N5cyIsImVjb3N5c3JjIiwib3BlbmlkIiwic2VuZGNvZGUiLCJzdmMiLCJvZmZsaW5lX2FjY2VzcyJdLCJhbXIiOlsiU21zQ29kZSJdfQ.f-S8C5nDs1H3_f0n6BDia-bmYmQ9EoQKhzJcNLqOBYHzv1jAxt042CiPg3-6IPqSyJEQKfPP0eTRUvE32jAgaVvVVu90NPHExa_9rdlCmgrV41B_HTcwOoOxG75zL8iD7EeU1KPJ-Gt4fGAm44LOl1U0tQOW_nJvmuaqy9pjTurejzMOIcLjptOvU6kX4XmavpmStSNmGZMNsoGXuDUxnSDEwFldSzIzci8JIRqhbCrF-1Jf9UoCVQ5y2S2QhKQsHm4F0-7A6Ts-88Xn6Jwpjw2Q9GbUzhpNA8uX7lblChOGdLV4E9xBPEoRe6TnnEMQicSeThr2EkBKjW_u9LC_aw")
		req.Header.Add("accessToken", user.Token)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		// 再设置个json
		//req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		log.Println(req.URL)
		log.Println(req.Header.Get("accessToken"))
		log.Println(req.PostForm)
		resp, err3 := client.Do(req)
		if err3 != nil {
			return
		}
		defer resp.Body.Close()
		out, err4 := ioutil.ReadAll(resp.Body)
		if err4 != nil {
			return
		}
		log.Println("response Body:", string(out))
		log.Println("response Status:", resp.Status)
		log.Println("response Headers:", resp.Header)
		duration := time.Duration(2) * time.Second
		time.Sleep(duration)
	}
	return
}
