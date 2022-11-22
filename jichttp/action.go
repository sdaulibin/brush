package jichttp

import (
	"binginx.com/brush/config"
	"binginx.com/brush/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func UserInfo(user *model.UserInfo, params *config.Params) (err error, jicUser *model.JicUser) {
	request, err := New(user, _userInfo_url, "GET", params)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(out, &jicUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	return err, jicUser
}

func Score(user *model.UserInfo, params *config.Params) (err error, score string) {
	request, err := New(user, _token_url, "GET", params)
	client := &http.Client{}
	resp, err := client.Do(request)
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

func News(user *model.UserInfo, params *config.Params) (err error, ecosysNewsId []model.Content) {
	request, err := New(user, _news_url, "GET", params)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
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
