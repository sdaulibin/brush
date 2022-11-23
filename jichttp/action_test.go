package jichttp

import (
	"binginx.com/brush/config"
	"binginx.com/brush/internal/logs"
	"binginx.com/brush/model"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	logs.Init()
	m.Run()
}

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
			"columnId": "ecosysNews",
			"pageNum":  "1",
			"pageSize": "200",
			"needAll":  "true",
		},
	}
	_, ecosysNews := News(&model.UserInfo{
		Token: _token,
	}, params)
	log.Println(len(ecosysNews))
	log.Println(ecosysNews)
}

func Test_Behavior(t *testing.T) {
	params := &config.Params{
		Params: map[string]string{
			"qtime":      strconv.Itoa(time.Now().Nanosecond()),
			"resourceId": "xxxxxx",
			"flag":       "1",
		},
	}
	err := DoBehavior(&model.UserInfo{
		Token: _token,
	}, params)
	log.Println(err)
}

func Test_View(t *testing.T) {
	params := &config.Params{
		Params: map[string]string{
			"qtime":     strconv.Itoa(time.Now().Nanosecond()),
			"contentId": "510268",
		},
	}
	err := View(&model.UserInfo{
		Token: _token,
	}, params)
	log.Println(err)
}

func Test_EcosysNews(t *testing.T) {
	err, ecosysNews := EcosysNews(model.UserInfo{
		Token: _token,
	})
	log.Println(ecosysNews)
	log.Println(err)
}
