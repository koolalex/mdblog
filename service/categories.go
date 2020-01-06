package service

import (
	"github.com/koolalex/mdblog/config"
	"github.com/koolalex/mdblog/models"
	"io/ioutil"
)

//GetCategories 获取目录
func GetCategories() (models.Categories, error) {
	var categories models.Categories
	categoriesDir, err := ioutil.ReadDir(config.Cfg.DocumentPath + "/content")
	if err != nil {
		return categories, err
	}

	categoryMap := make(map[string]*models.Category)
	for _, category := range categoriesDir {
		if !category.IsDir() {
			continue
		}

		if markdownList, err := models.GetMarkdownListByCache("/" + category.Name()); err == nil {
			for _, md := range markdownList {
				category, exists := categoryMap[md.Category]
				if !exists {
					category = &models.Category{}
					categoryMap[md.Category] = category
				}
				category.Name = md.Meta.Category
				category.MarkdownFileList = append(category.MarkdownFileList, md)
			}
		} else {
			return categories, err
		}
	}

	for _, category := range categoryMap {
		markdownList := category.MarkdownFileList
		listLen := len(markdownList)
		categoryListFileNumber := listLen
		if listLen >= config.Cfg.CategoryListFileNumber {
			categoryListFileNumber = config.Cfg.CategoryListFileNumber
		}

		category.Number = listLen
		category.MarkdownFileList = markdownList[0:categoryListFileNumber]
		categories = append(categories, *category)
	}

	return categories, nil
}

func GetCategory(categoryName string) (models.Category, error) {
	category := models.Category{
		Name: categoryName,
	}

	categoriesDir, err := ioutil.ReadDir(config.Cfg.DocumentPath + "/content")
	if err != nil {
		return category, err
	}

	for _, fi := range categoriesDir {
		if !fi.IsDir() {
			continue
		}

		if markdownList, err := models.GetMarkdownListByCache("/" + fi.Name()); err == nil {
			for _, md := range markdownList {
				if md.Meta.Category == categoryName {
					category.Number += 1
					category.MarkdownFileList = append(category.MarkdownFileList, md)
				}
			}
		} else {
			return category, err
		}
	}

	return category, nil
}
