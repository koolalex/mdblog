package controller

import (
	"github.com/koolalex/mdblog/config"
	"github.com/koolalex/mdblog/models"
	"github.com/koolalex/mdblog/service"
	"net/http"
	"strconv"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		page = 1
	}

	template, err := HtmlTemplate("index")
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	searchKey := r.Form.Get("search")
	markdownPagination, err := service.GetArticleList(page, "/", searchKey)
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "首页",
		"Data":   markdownPagination,
		"Config": config.Cfg,
	})
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}
}

func Categories(w http.ResponseWriter, r *http.Request) {
	template, err := HtmlTemplate("categories")
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	categories, err := service.GetCategories()
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "分类",
		"Data":   categories,
		"Config": config.Cfg,
	})
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}
}

func Works(w http.ResponseWriter, r *http.Request) {
	template, err := HtmlTemplate("works")
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	markdown, err := models.ReadMarkdownBody("/Works.md")
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "作品",
		"Data":   markdown,
		"Config": config.Cfg,
	})
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	template, err := HtmlTemplate("about")
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	markdown, err := models.ReadMarkdownBody("/About.md")
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "关于",
		"Data":   markdown,
		"Config": config.Cfg,
	})
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}
}

/*sub page*/
func Article(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	path := r.Form.Get("path")
	template, err := HtmlTemplate("article")
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	article, err := models.GetMarkdownDetails(path)
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "文章详情",
		"Data":   article,
		"Config": config.Cfg,
	})
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}
}

func CategoryArticle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	template, err := HtmlTemplate("category")
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	categoryName := r.Form.Get("name")
	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		page = 1
	}
	content, err := service.GetCategoryArticlePagination(page, categoryName, "")
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  strings.Replace(categoryName, "/", "", 1),
		"Data":   content,
		"Config": config.Cfg,
	})
	if err != nil {
		WriteErrorHtml(w, err.Error())
		return
	}
}
