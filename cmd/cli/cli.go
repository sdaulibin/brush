package main

import (
	"binginx.com/brush/cmd/api"
	"binginx.com/brush/config"
	"binginx.com/brush/internal/clients"
	"binginx.com/brush/internal/logs"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	config.MustInit()
	logs.Init()
	clients.MustInit()
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "启动brush",
			Action: func(c *cli.Context) error {
				api.Run()
				return nil
			},
		},
		{
			Name:    "token",
			Aliases: []string{"a"},
			Usage:   "直接使用token自动增加分数",
			Action: func(context *cli.Context) error {
				token := context.Args().Get(0)
				if token != "" {
					api.Total(token)
					score := api.Score(token)
					logs.Logger.Infof("当前分数:%v", score)
				}
				return nil
			},
		},
		{
			Name:    "info",
			Aliases: []string{"i"},
			Usage:   "获取当前分数",
			Action: func(context *cli.Context) error {
				token := context.Args().Get(0)
				if token != "" {
					score := api.Score(token)
					logs.Logger.Infof("当前分数:%v", score)
				}
				return nil
			},
		},
		{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "使用文件进行分数增加",
			Action: func(context *cli.Context) error {
				path := context.Args().Get(0)
				if path != "" {
					api.Import(path)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
