package api

import (
	"binginx.com/brush/config"
	"binginx.com/brush/internal/logs"
	"binginx.com/brush/model"
	"binginx.com/brush/service"
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func getNews(columnId, token string) []model.Content {
	newsParams := &config.Params{
		Params: map[string]string{
			"qtime":    strconv.Itoa(time.Now().Nanosecond()),
			"columnId": columnId,
			"pageNum":  "1",
			"pageSize": "300",
			"needAll":  "true",
		},
	}
	err, news := service.News(&model.UserInfo{Token: token}, newsParams)
	if err != nil {
		log.Println(err)
	}
	return news
}

func Total(token string) {
	contents := []model.Content{}
	unEcosysNews := getNews("unEcosysNews", token)
	contents = append(contents, unEcosysNews...)
	ecosysNews := getNews("ecosysNews", token)
	contents = append(contents, ecosysNews...)
	companyNews := getNews("companyNews", token)
	contents = append(contents, companyNews...)
	for _, news := range contents {
		viewParams := &config.Params{
			Params: map[string]string{
				"qtime":     strconv.Itoa(time.Now().Nanosecond()),
				"contentId": strconv.Itoa(news.ContentId),
			},
		}
		service.View(&model.UserInfo{Token: token}, viewParams)
		params := &config.Params{
			Params: map[string]string{
				"qtime":        strconv.Itoa(time.Now().Nanosecond()),
				"resourceId":   strconv.Itoa(news.ContentId),
				"userBehavior": service.Behavior_COLLECTION,
			},
		}
		service.DoBehavior(&model.UserInfo{Token: token}, params)
		params = &config.Params{
			Params: map[string]string{
				"qtime":        strconv.Itoa(time.Now().Nanosecond()),
				"resourceId":   strconv.Itoa(news.ContentId),
				"userBehavior": service.Behavior_PRAISE,
			},
		}
		service.DoBehavior(&model.UserInfo{Token: token}, params)
	}
}

func Score(token string) string {
	err, score := service.Score(&model.UserInfo{
		Token: token,
	}, &config.Params{
		Params: map[string]string{
			"qtime": strconv.Itoa(time.Now().Nanosecond()),
		},
	})
	if err != nil {
		log.Println(err)
	}
	return score
}

func Import(filePath string) {
	var tokens = []string{}
	fi, e := os.Open(filePath)
	if e != nil {
		logs.Logger.Infof("read file error:%v", e.Error())
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		tokens = append(tokens, string(a))
	}
	for i := 0; i < len(tokens); i++ {
		Total(tokens[i])
	}
}
