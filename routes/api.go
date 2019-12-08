package routes

import (
	"ForestBlog/config"
	"ForestBlog/controller"
	"net/http"
)

func initApiRoute() {

	http.HandleFunc(config.Cfg.GitHookUrl, controller.GithubHook)

}
