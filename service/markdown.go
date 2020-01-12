package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/koolalex/mdblog/config"
	"github.com/koolalex/mdblog/library/cache"
	"github.com/koolalex/mdblog/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func readMarkdown(path string) (models.Markdown, models.MarkdownDetails, error) {
	var (
		markdownDoc    models.Markdown
		markdownDetail models.MarkdownDetails
	)

	fullPath := config.Cfg.DocumentPath + "/content" + path
	if val, exists := Cache.Get(fullPath); exists {
		mdd := val.(models.MarkdownDetails)
		md := mdd.Markdown
		fmt.Println("read from cache title: " + md.Title)
		return md, mdd, nil
	}

	meta, err := readMeta(path)
	if err != nil {
		return markdownDoc, markdownDetail, err
	}

	markdownFile, fileErr := os.Stat(fullPath)
	if fileErr != nil {
		return markdownDoc, markdownDetail, fileErr
	}
	if markdownFile.IsDir() {
		return markdownDoc, markdownDetail, errors.New("this path is Dir")
	}

	markdownBytes, mdErr := ioutil.ReadFile(fullPath)
	if mdErr != nil {
		return markdownDoc, markdownDetail, mdErr
	}

	markdownBytes = bytes.TrimSpace(markdownBytes)
	markdownDoc.Path = path
	markdownDoc.Category = meta.Category
	markdownDoc.Title = meta.Title
	markdownDoc.Date = models.JsonTime(markdownFile.ModTime())

	markdownDetail.Markdown = markdownDoc
	markdownDetail.Body = string(markdownBytes)

	if !bytes.HasPrefix(markdownBytes, []byte("```json")) {
		markdownDoc.Description = cropDesc(markdownBytes)
		Cache.SetDefault(fullPath, markdownDetail)
		return markdownDoc, markdownDetail, nil
	}

	markdownBytes = bytes.Replace(markdownBytes, []byte("```json"), []byte(""), 1)
	markdownArrInfo := bytes.SplitN(markdownBytes, []byte("```"), 2)

	markdownDoc.Description = cropDesc(markdownArrInfo[1])
	if err := json.Unmarshal(bytes.TrimSpace(markdownArrInfo[0]), &markdownDoc); err != nil {
		return markdownDoc, markdownDetail, err
	}

	markdownDoc.Path = path //保证Path不被用户json赋值，json不能添加`json:"-"`忽略，否则编码到缓存的时候会被忽悠。
	markdownDetail.Markdown = markdownDoc
	markdownDetail.Body = string(markdownArrInfo[1])
	Cache.SetDefault(fullPath, markdownDetail)
	return markdownDoc, markdownDetail, nil
}

func readMeta(path string) (meta models.Meta, err error) {
	categoryName := strings.Replace(path, "/", "", 1)
	if strings.Index(categoryName, "/") >= 0 { //文件在根目录下(content/)没有分类名称
		categoryName = strings.Split(categoryName, "/")[0]
		metaFullPath := config.Cfg.DocumentPath + "/content/" + categoryName + "/meta.yml"
		data, err := ioutil.ReadFile(metaFullPath)
		if err != nil {
			return models.Meta{}, err
		}
		if err = yaml.Unmarshal([]byte(data), &meta); err == nil {
			return meta, nil
		}
	}
	return models.Meta{}, errors.New(fmt.Sprintf("path invalid:%v", path))
}

func cropDesc(c []byte) string {
	content := []rune(string(c))
	contentLen := len(content)

	if contentLen <= config.Cfg.DescriptionLen {
		return string(content[0:contentLen])
	}

	return string(content[0:config.Cfg.DescriptionLen])
}

//GetMarkdown 读取路径下的md文件的部分信息json
func GetMarkdown(path string) (models.Markdown, error) {
	if content, _, err := readMarkdown(path); err != nil {
		return content, err
	} else {
		return content, nil
	}
}

//读取路径下的md文件完整信息
func GetMarkdownDetails(path string) (models.MarkdownDetails, error) {
	_, content, err := readMarkdown(path)
	if err != nil {
		return content, err
	}
	return content, nil
}

//递归获取md文件信息
func getMarkdownList(dir string) (models.MarkdownList, error) {
	var mdList models.MarkdownList
	fullDir := config.Cfg.DocumentPath + "/content" + dir
	fileOrDir, err := ioutil.ReadDir(fullDir)
	if err != nil {
		return mdList, err
	}

	for _, fileInfo := range fileOrDir {
		var subDir string
		if "/" == dir {
			subDir = "/" + fileInfo.Name()
		} else {
			subDir = dir + "/" + fileInfo.Name()
		}
		if fileInfo.IsDir() {
			subMdList, err := getMarkdownList(subDir)
			if err != nil {
				return mdList, err
			}
			mdList = append(mdList, subMdList...)
		} else if strings.HasSuffix(strings.ToLower(fileInfo.Name()), "md") {
			if markdown, err := GetMarkdown(subDir); err != nil {
				return mdList, err
			} else {
				mdList = append(mdList, markdown)
			}
		}
	}
	return mdList, nil
}

func GetMarkdownListByCache(dir string) (models.MarkdownList, error) {
	var markdownLists models.MarkdownList
	if docs, exists := Cache.Get(dir); exists {
		//fmt.Println("Get MarkdownList From Cache..")
		return docs.(models.MarkdownList), nil
	}

	markdownLists, err := getMarkdownList(dir)
	if err != nil {
		return markdownLists, err
	}

	sort.Sort(markdownLists)

	Cache.Set(dir, markdownLists, cache.NoExpiration)

	return markdownLists, nil
}

func ReadMarkdownBody(path string) (string, error) {
	fullPath := config.Cfg.DocumentPath + path
	markdown, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return "", err
	}

	return string(markdown), nil
}
