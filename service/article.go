package service

import (
	"fmt"
	"github.com/koolalex/mdblog/config"
	"github.com/koolalex/mdblog/models"
	"log"
	"math"
	"os/exec"
	"strings"
)

func GetArticleList(page int, dir string, search string) (models.MarkdownPagination, error) {
	allArticle, err := GetMarkdownListByCache(dir)
	if err != nil {
		return models.MarkdownPagination{}, err
	}
	if "" == search {
		return getPaginationData(allArticle, page)
	}

	var newArticleList models.MarkdownList
	for _, article := range allArticle {
		if strings.Index(article.Title, search) != -1 {
			newArticleList = append(newArticleList, article)
		}
	}
	return getPaginationData(newArticleList, page)
}

func GetCategoryArticlePagination(page int, categoryName string, search string) (models.MarkdownPagination, error) {
	category, err := GetCategory(categoryName)
	if err != nil {
		return models.MarkdownPagination{}, err
	}
	if "" == search {
		return getPaginationData(category.MarkdownFileList, page)
	}

	var newArticleList models.MarkdownList
	for _, article := range category.MarkdownFileList {
		if strings.Index(article.Title, search) != -1 {
			newArticleList = append(newArticleList, article)
		}
	}
	return getPaginationData(newArticleList, page)
}

func UpdateArticle() {
	blogPath := config.CurrentDir + "/" + config.Cfg.DocumentPath

	_, err := exec.LookPath("git")
	if err != nil {
		fmt.Println("请先安装git并克隆博客文档到" + blogPath)
		log.Fatalf("git cmd failed with %s\n", err)
	}

	cmd := exec.Command("git", "pull")
	cmd.Dir = blogPath
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	log.Println("UpdateArticle:" + string(out))
	Cache.Delete("docs")
	_, err = GetMarkdownListByCache("/") //生成缓存
	if err != nil {
		log.Fatalf("生成缓存失败： %s\n", err)
	}
}

func getPaginationData(allArticle models.MarkdownList, page int) (models.MarkdownPagination, error) {
	var paginationData models.MarkdownPagination
	articleLen := len(allArticle)
	pageSize := config.Cfg.PageSize
	totalPage := int(math.Floor(float64(articleLen / pageSize)))
	if (articleLen % pageSize) != 0 {
		totalPage++
	}

	paginationData.Total = articleLen
	paginationData.CurrentPage = page
	paginationData.PageNumber = buildArrByInt(totalPage)
	if page < 1 || pageSize*(page-1) > articleLen { //超出页码

		paginationData.CurrentPage = 1

		if pageSize <= articleLen {
			paginationData.Markdowns = allArticle[0:pageSize]
		} else {
			paginationData.Markdowns = allArticle[0:articleLen]
		}
		return paginationData, nil
	}

	startNum := (page - 1) * pageSize
	endNum := startNum + pageSize
	if endNum > articleLen {
		paginationData.Markdowns = allArticle[startNum:articleLen]
	} else {
		paginationData.Markdowns = allArticle[startNum:endNum]
	}
	return paginationData, nil
}

func buildArrByInt(num int) []int {
	var arr []int
	for i := 1; i <= num; i++ {
		arr = append(arr, i)
	}
	return arr
}
