package jichttp

import (
	"binginx.com/brush/config"
	"binginx.com/brush/model"
	"log"
	"strconv"
	"testing"
	"time"
)

func Test_UserInfo(t *testing.T) {

	err, userInfo := UserInfo(&model.UserInfo{
		Token: _token,
	}, &config.Params{
		Params: map[string]string{
			"qtime": strconv.Itoa(time.Now().Nanosecond()),
		},
	})
	log.Println(userInfo)
	log.Println(err)
}

func Test_Score(t *testing.T) {

	err, score := Score(&model.UserInfo{
		Token: _token,
	}, &config.Params{
		Params: map[string]string{
			"qtime": strconv.Itoa(time.Now().Nanosecond()),
		},
	})
	log.Println(score)
	log.Println(err)
}
func Test_News(t *testing.T) {
	params := &config.Params{
		Params: map[string]string{
			"qtime":    strconv.Itoa(time.Now().Nanosecond()),
			"columnId": "unEcosysNews",
			"pageNum":  "1",
			"pageSize": "500",
			"needAll":  "true",
		},
	}
	err, ecosysNews := News(&model.UserInfo{
		Token: _token,
	}, params)
	log.Println(len(ecosysNews))
	log.Println(err)
}

func Test_EcosysNews(t *testing.T) {
	err, ecosysNews := EcosysNews(model.UserInfo{
		Token: _token,
	})
	log.Println(ecosysNews)
	log.Println(err)
}
