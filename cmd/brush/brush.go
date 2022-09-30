package main

import (
	"binginx.com/brush/jichttp"
	"binginx.com/brush/model"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	defaultConfig = "/etc/brush/config.ini"
)

func main() {
	runServer("127.0.0.1:8080")
}

func runServer(addr string) {
	var tokens = []string{}
	fi, e := os.Open("/Users/binginx/Downloads/token2.txt")
	if e != nil {
		fmt.Println("read file error")
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		//fmt.Println(string(a))
		tokens = append(tokens, string(a))
	}
	fmt.Println(cap(tokens))
	router := gin.Default()
	router.GET("/brush/token", func(c *gin.Context) {
		start := time.Now()
		var wg sync.WaitGroup
		wg.Add(len(tokens))
		for i := 0; i < len(tokens); i++ {
			//defer wg.Done()
			//go token(c, tokens[i])
			token(c, tokens[i])
		}
		//wg.Wait()
		elapsed := time.Since(start)
		fmt.Printf("WaitGroupStart Time %s\n ", elapsed)
	})
	router.GET("/brush/action", func(c *gin.Context) {
		start := time.Now()
		var wg sync.WaitGroup
		wg.Add(len(tokens))
		for i := 0; i < len(tokens); i++ {
			defer wg.Done()
			go action(c, tokens[i])
			//action(c, tokens[i])
		}
		wg.Wait()
		elapsed := time.Since(start)
		fmt.Printf("WaitGroupStart Time %s\n ", elapsed)
	})
	router.Run(addr) // 默认监听并在 0.0.0.0:8080 上启动服务
}

func token(c *gin.Context, token string) {
	var user model.UserInfo = model.UserInfo{
		UserName: "xxx",
		Token:    token,
	}
	err, score := jichttp.ScoreTotal(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "successful!=====" + score})
}

func action(c *gin.Context, token string) {
	var user model.UserInfo = model.UserInfo{
		UserName: "xxx",
		Token:    token,
	}
	//err := c.BindJSON(&user)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//log.Println("user:", user.UserName, user.Token)
	err0, news := jichttp.EcosysNews(user)
	if err0 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err0.Error()})
	}
	err := jichttp.ViewNews(user, news)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err1 := jichttp.Behavior(user, "COLLECTION", news)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
	}
	err2 := jichttp.Behavior(user, "PRAISE", news)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "successful!"})
}
