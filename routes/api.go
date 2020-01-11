package routes

import (
	"github.com/koolalex/mdblog/config"
	"github.com/koolalex/mdblog/controller"
	"net/http"
)

func initApiRoute() {
	http.HandleFunc(config.Cfg.GitHookUrl, controller.GithubHook)
}
