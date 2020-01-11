package main

import (
	"fmt"
	"github.com/koolalex/mdblog/config"
	"github.com/koolalex/mdblog/routes"
	"github.com/koolalex/mdblog/service"
	"net/http"
	"strconv"
)

func main() {
	//load
	routes.InitRoute()

	fmt.Printf("App Name： %v \n", config.Cfg.AppName)
	fmt.Printf("Markdown Blog Version: v%v \n", config.Cfg.Version)
	fmt.Printf("Listen On Port: %v \n", config.Cfg.Port)
	fmt.Printf("Posts Update GitHookUrl: %v   WebHookSecret:%v \n", config.Cfg.GitHookUrl, config.Cfg.WebHookSecret)
	service.UpdateArticle()

	if err := http.ListenAndServe(":"+strconv.Itoa(config.Cfg.Port), nil); err != nil {
		fmt.Println("Listen Error:", err)
	}

	c := make(chan bool)
	<-c
}
