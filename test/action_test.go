package test

import (
	"binginx.com/brush/config"
	"binginx.com/brush/internal/logs"
	"binginx.com/brush/model"
	"binginx.com/brush/service"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	logs.Init()
	m.Run()
}

func Test_UserInfo(t *testing.T) {

	err, userInfo := service.UserInfo(&model.UserInfo{
		Token: service.Token,
	}, &config.Params{
		Params: map[string]string{
			"qtime": strconv.Itoa(time.Now().Nanosecond()),
		},
	})
	log.Println(userInfo)
	log.Println(err)
}

func Test_Score(t *testing.T) {
	err, score := service.Score(&model.UserInfo{
		Token: service.Token,
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
	_, ecosysNews := service.News(&model.UserInfo{
		Token: service.Token,
	}, params)
	log.Println(len(ecosysNews))
	log.Println(ecosysNews)
}

func Test_Behavior(t *testing.T) {
	params := &config.Params{
		Params: map[string]string{
			"qtime":      strconv.Itoa(time.Now().Nanosecond()),
			"resourceId": "xxxxxx",
			"flag":       service.Behavior_COLLECTION,
		},
	}
	err := service.DoBehavior(&model.UserInfo{
		Token: service.Token,
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
	err := service.View(&model.UserInfo{
		Token: service.Token,
	}, params)
	log.Println(err)
}

func Test_EcosysNews(t *testing.T) {
	err, ecosysNews := service.EcosysNews(model.UserInfo{
		Token: service.Token,
	})
	log.Println(ecosysNews)
	log.Println(err)
}

func TestTotal(t *testing.T) {
	newsParams := &config.Params{
		Params: map[string]string{
			"qtime":    strconv.Itoa(time.Now().Nanosecond()),
			"columnId": "ecosysNews",
			"pageNum":  "1",
			"pageSize": "300",
			"needAll":  "true",
		},
	}
	err, ecosysNews := service.News(&model.UserInfo{Token: service.Token}, newsParams)
	if err != nil {
		t.Log(err)
	}
	for _, news := range ecosysNews {
		viewParams := &config.Params{
			Params: map[string]string{
				"qtime":     strconv.Itoa(time.Now().Nanosecond()),
				"contentId": strconv.Itoa(news.ContentId),
			},
		}
		service.View(&model.UserInfo{Token: service.Token}, viewParams)
		params := &config.Params{
			Params: map[string]string{
				"qtime":        strconv.Itoa(time.Now().Nanosecond()),
				"resourceId":   strconv.Itoa(news.ContentId),
				"userBehavior": service.Behavior_COLLECTION,
			},
		}
		service.DoBehavior(&model.UserInfo{Token: service.Token}, params)
		params = &config.Params{
			Params: map[string]string{
				"qtime":        strconv.Itoa(time.Now().Nanosecond()),
				"resourceId":   strconv.Itoa(news.ContentId),
				"userBehavior": service.Behavior_PRAISE,
			},
		}
		service.DoBehavior(&model.UserInfo{Token: service.Token}, params)
	}
}

func TestImportToken(t *testing.T) {
	var tokens = []string{}
	fi, e := os.Open("/Users/binginx/Downloads/token.txt")
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
		newsParams := &config.Params{
			Params: map[string]string{
				"qtime":    strconv.Itoa(time.Now().Nanosecond()),
				"columnId": "ecosysNews",
				"pageNum":  "1",
				"pageSize": "200",
				"needAll":  "true",
			},
		}
		err, ecosysNews := service.News(&model.UserInfo{Token: service.Token}, newsParams)
		if err != nil {
			t.Log(err)
		}
		idChan := make(chan string, len(ecosysNews))
		errChan := make(chan error, len(ecosysNews))
		for _, news := range ecosysNews {
			go func(content *model.Content) {
				viewParams := &config.Params{
					Params: map[string]string{
						"qtime":     strconv.Itoa(time.Now().Nanosecond()),
						"contentId": strconv.Itoa(news.ContentId),
					},
				}
				err = service.View(&model.UserInfo{Token: service.Token}, viewParams)
				if err != nil {
					errChan <- fmt.Errorf("view fail:%v", err.Error())
					return
				}
				idChan <- strconv.Itoa(news.ContentId) + ":view"
				params := &config.Params{
					Params: map[string]string{
						"qtime":        strconv.Itoa(time.Now().Nanosecond()),
						"resourceId":   strconv.Itoa(news.ContentId),
						"userBehavior": service.Behavior_COLLECTION,
					},
				}
				err = service.DoBehavior(&model.UserInfo{Token: service.Token}, params)
				if err != nil {
					errChan <- fmt.Errorf("behavior collection fail:%v", err.Error())
					return
				}
				idChan <- strconv.Itoa(news.ContentId) + ":collection"
				params = &config.Params{
					Params: map[string]string{
						"qtime":        strconv.Itoa(time.Now().Nanosecond()),
						"resourceId":   strconv.Itoa(news.ContentId),
						"userBehavior": service.Behavior_PRAISE,
					},
				}
				err = service.DoBehavior(&model.UserInfo{Token: service.Token}, params)
				if err != nil {
					errChan <- fmt.Errorf("behavior praise fail:%v", err.Error())
					return
				}
				idChan <- strconv.Itoa(news.ContentId) + ":praise"
			}(&news)
		}
		for i := 0; i < len(ecosysNews); i++ {
			select {
			case err = <-errChan:
			case idDesc := <-idChan:
				logs.Logger.Infoln(idDesc)
			}
		}
	}
}
