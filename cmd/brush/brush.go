package main

import (
	"binginx.com/brush/jichttp"
	"binginx.com/brush/model"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"net/http"
	"os"
)

//var (
//	h     bool
//	token string
//	score string
//)
var (
	installCommand = flag.NewFlagSet("install", flag.ExitOnError)
	runCommand     = flag.NewFlagSet("run", flag.ExitOnError)
	user           string
	workingdir     string
	config         string
)

const (
	defaultConfig = "/etc/brush/config.ini"
)

//func init() {
//	flag.BoolVar(&h, "h", false, "help")
//	flag.StringVar(&token, "t", "", "中投生态圈token")
//	flag.StringVar(&token, "s", "", "中投生态圈分数")
//}
func usage() {
	s := `
USAGE:
   brush command [command options] 


COMMANDS:
     install               install service
     uninstall             uninstall service
     start                 start service
     stop                  stop service
	 run                   run service

OPTIONS:
     -config string
       	config file of the service (default "/etc/brush/config.ini")
     -user string
       	user account to run the service
     -workingdir string
    	working directory of the service`

	fmt.Println(s)
}

func init() {
	installCommand.StringVar(&user, "user", "", "user account to run the service")
	installCommand.StringVar(&workingdir, "workingdir", "", "working directory of the service")
	installCommand.StringVar(&config, "config", "/etc/brush/config.ini", "config file of the service")
	runCommand.StringVar(&config, "config", defaultConfig, "config file of the service")
	flag.Usage = usage
}

func main() {
	//flag.Parse()
	//
	//if h {
	//	flag.Usage()
	//	return
	//}
	//
	//if token == "" {
	//	flag.Usage()
	//	return
	//}

	var err error
	n := len(os.Args)
	if n <= 1 {
		fmt.Printf("invalid args\n")
		flag.Usage()
		return
	}

	subCmd := os.Args[1] // the second arg

	// get Config
	c, err := getServiceConfig(subCmd)
	if err != nil {
		fmt.Printf("get service config error: %s\n", err)
		return
	}

	prg := &NullService{}
	srv, err := service.New(prg, c)
	if err != nil {
		fmt.Printf("new service error: %s\n", err)
		return
	}

	err = runServiceControl(srv, subCmd)
	if err != nil {
		fmt.Printf("%s operation error: %s\n", subCmd, err)
		return
	}

	fmt.Printf("%s operation ok\n", subCmd)
	return
}

func runServiceControl(srv service.Service, subCmd string) error {
	switch subCmd {
	case "run":
		return run(config)
	default:
		return service.Control(srv, subCmd)
	}
}

func run(config string) error {
	// load info from config
	c, err := loadConfigFromIni(config)
	if err != nil {
		return err
	}

	runServer(c.Server.Addr)
	return err
}

type NullService struct{}

func (p *NullService) Start(s service.Service) error {
	return nil
}

func (p *NullService) Stop(s service.Service) error {
	return nil
}

func runServer(addr string) {
	router := gin.Default()
	router.GET("/brush/token", func(c *gin.Context) {
		var user model.UserInfo = model.UserInfo{
			UserName: "libin",
			Token:    "eyJhbGciOiJSUzI1NiIsInR5cCI6ImF0K2p3dCJ9.eyJuYmYiOjE2NjE5OTU2MjcsImV4cCI6MTY2MTk5OTIyNywiaXNzIjoiaHR0cHM6Ly9pZHMuY2hpbmEtaW52LmNuIiwiYXVkIjpbImVjb3N5cyIsImVjb3N5c3JjIiwic2VuZGNvZGUiLCJzdmMiXSwiY2xpZW50X2lkIjoiZWNvc3lzIiwic3ViIjoibGliaW5Ac2RqaWN0ZWMuY29tIiwiYXV0aF90aW1lIjoxNjYxOTk1NjI3LCJpZHAiOiJsb2NhbCIsIm5hbWUiOiLmnY7lvawiLCJlbWFpbCI6ImxpYmluQHNkamljdGVjLmNvbSIsImRlcGFydG1lbnQiOiLkuLTml7bpg6jpl6giLCJjb21wYW55Ijoi5bu65oqV5pWw5o2u56eR5oqA77yI5bGx5Lic77yJ5pyJ6ZmQ5YWs5Y-4IiwiZGl2aXNpb24iOiIiLCJlbXBsb3llZWlkIjoiMTIzODk1IiwiZW1haWxfdmVyaWZpZWQiOiJ0cnVlIiwianRpIjoiQkRCRDBGNkI4QTQxNDk5NUU4NTUxNjY4RTM0NkRBN0IiLCJpYXQiOjE2NjE5OTU2MjcsInNjb3BlIjpbImVjb3N5cyIsImVjb3N5c3JjIiwib3BlbmlkIiwic2VuZGNvZGUiLCJzdmMiLCJvZmZsaW5lX2FjY2VzcyJdLCJhbXIiOlsiU21zQ29kZSJdfQ.f-S8C5nDs1H3_f0n6BDia-bmYmQ9EoQKhzJcNLqOBYHzv1jAxt042CiPg3-6IPqSyJEQKfPP0eTRUvE32jAgaVvVVu90NPHExa_9rdlCmgrV41B_HTcwOoOxG75zL8iD7EeU1KPJ-Gt4fGAm44LOl1U0tQOW_nJvmuaqy9pjTurejzMOIcLjptOvU6kX4XmavpmStSNmGZMNsoGXuDUxnSDEwFldSzIzci8JIRqhbCrF-1Jf9UoCVQ5y2S2QhKQsHm4F0-7A6Ts-88Xn6Jwpjw2Q9GbUzhpNA8uX7lblChOGdLV4E9xBPEoRe6TnnEMQicSeThr2EkBKjW_u9LC_aw",
		}
		//err := c.BindJSON(&user)
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//	return
		//}
		//log.Println("user:", user.UserName, user.Token)
		err, score := jichttp.ScoreTotal(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "successful!=====" + score})
	})
	router.GET("/brush/action", func(c *gin.Context) {
		var user model.UserInfo = model.UserInfo{
			UserName: "libin",
			Token:    "eyJhbGciOiJSUzI1NiIsInR5cCI6ImF0K2p3dCJ9.eyJuYmYiOjE2NjE5OTkzMzUsImV4cCI6MTY2MjAwMjkzNSwiaXNzIjoiaHR0cHM6Ly9pZHMuY2hpbmEtaW52LmNuIiwiYXVkIjpbImVjb3N5cyIsImVjb3N5c3JjIiwic2VuZGNvZGUiLCJzdmMiXSwiY2xpZW50X2lkIjoiZWNvc3lzIiwic3ViIjoibGliaW5Ac2RqaWN0ZWMuY29tIiwiYXV0aF90aW1lIjoxNjYxOTk1NjI3LCJpZHAiOiJsb2NhbCIsIm5hbWUiOiLmnY7lvawiLCJlbWFpbCI6ImxpYmluQHNkamljdGVjLmNvbSIsImRlcGFydG1lbnQiOiLkuLTml7bpg6jpl6giLCJjb21wYW55Ijoi5bu65oqV5pWw5o2u56eR5oqA77yI5bGx5Lic77yJ5pyJ6ZmQ5YWs5Y-4IiwiZGl2aXNpb24iOiIiLCJlbXBsb3llZWlkIjoiMTIzODk1IiwiZW1haWxfdmVyaWZpZWQiOiJ0cnVlIiwianRpIjoiQkRCRDBGNkI4QTQxNDk5NUU4NTUxNjY4RTM0NkRBN0IiLCJpYXQiOjE2NjE5OTU2MjcsInNjb3BlIjpbImVjb3N5cyIsImVjb3N5c3JjIiwib3BlbmlkIiwic2VuZGNvZGUiLCJzdmMiLCJvZmZsaW5lX2FjY2VzcyJdLCJhbXIiOlsiU21zQ29kZSJdfQ.EaVhU5W0Hkxtok0uGC9XHL6O11hBE5oKPHw5n5lFOXw4qeBONUv-nC3jv84fkdnJrjqFr7tl6xt1HNEhfFevdq2FGRlCDybSwfkLGBjl0tTSnmoYMVHYPmvurGKkMC4YjC-ry747E-OTEZiZy1KdQnQVJxiqrNxPGG90hlH0_7kjpt--VLZ7npsus5JoXbuD3Olx5ZCG46B5pu9bgAxwtRbEiQezlGDbxyx3y-q5d4WocSu5L4x5iCZNBet8xWrRCD-3lBvqioJU5MNPS0zafyWXgQ190syp6g8j40EswxsXywkXC79O1EDyIawrZfa6oHrRodMAkQb-wFwn7-ZU5w",
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
	})
	router.Run(addr) // 默认监听并在 0.0.0.0:8080 上启动服务
}
