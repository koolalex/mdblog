package routes

import (
	"net/http"

	"github.com/koolalex/mdblog/controller"
)

func initWebRoute() {
	//一级页面
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/categories", controller.Categories)
	http.HandleFunc("/works", controller.Works)
	http.HandleFunc("/about", controller.About)
	//二级页面
	http.HandleFunc("/article", controller.Article)
	http.HandleFunc("/category", controller.CategoryArticle)
	//静态文件服务器
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("resources/public"))))
	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(config.Cfg.DocumentPath+"/assets"))))
	http.HandleFunc("/assets/", controller.ServAssets)

}
