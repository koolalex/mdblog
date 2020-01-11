package models

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/koolalex/mdblog/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func readMarkdown(path string) (Markdown, MarkdownDetails, error) {
	var (
		content     Markdown
		fullContent MarkdownDetails
	)

	fullPath := config.Cfg.DocumentPath + "/content" + path
	meta, err := readMeta(path)
	if err != nil {
		return content, fullContent, err
	}

	markdownFile, fileErr := os.Stat(fullPath)

	if fileErr != nil {
		return content, fullContent, fileErr
	}
	if markdownFile.IsDir() {
		return content, fullContent, errors.New("this path is Dir")
	}
	markdown, mdErr := ioutil.ReadFile(fullPath)

	if mdErr != nil {
		return content, fullContent, mdErr
	}
	markdown = bytes.TrimSpace(markdown)

	content.Path = path
	content.Category = meta.Category
	content.Title = meta.Title
	content.Date = JsonTime(markdownFile.ModTime())

	fullContent.Markdown = content
	fullContent.Body = string(markdown)

	if !bytes.HasPrefix(markdown, []byte("```json")) {
		content.Description = cropDesc(markdown)
		return content, fullContent, nil
	}

	markdown = bytes.Replace(markdown, []byte("```json"), []byte(""), 1)
	markdownArrInfo := bytes.SplitN(markdown, []byte("```"), 2)

	content.Description = cropDesc(markdownArrInfo[1])

	if err := json.Unmarshal(bytes.TrimSpace(markdownArrInfo[0]), &content); err != nil {
		return content, fullContent, err
	}

	content.Path = path //保证Path不被用户json赋值，json不能添加`json:"-"`忽略，否则编码到缓存的时候会被忽悠。
	fullContent.Markdown = content
	fullContent.Body = string(markdownArrInfo[1])

	return content, fullContent, nil
}

func readMeta(path string) (meta Meta, err error) {
	categoryName := strings.Replace(path, "/", "", 1)
	if strings.Index(categoryName, "/") >= 0 { //文件在根目录下(content/)没有分类名称
		categoryName = strings.Split(categoryName, "/")[0]
		metaFullPath := config.Cfg.DocumentPath + "/content/" + categoryName + "/meta.yml"
		data, err := ioutil.ReadFile(metaFullPath)
		if err != nil {
			return Meta{}, err
		}
		if err = yaml.Unmarshal([]byte(data), &meta); err == nil {
			return meta, nil
		}
	}
	return Meta{}, errors.New(fmt.Sprintf("path invalid:%v", path))
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
func GetMarkdown(path string) (Markdown, error) {
	if content, _, err := readMarkdown(path); err != nil {
		return content, err
	} else {
		return content, nil
	}
}

//读取路径下的md文件完整信息
func GetMarkdownDetails(path string) (MarkdownDetails, error) {
	_, content, err := readMarkdown(path)

	if err != nil {
		return content, err
	}

	return content, nil
}

//递归获取md文件信息
func getMarkdownList(dir string) (MarkdownList, error) {
	var mdList MarkdownList
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

func GetMarkdownListByCache(dir string) (MarkdownList, error) {
	cacheFileName := fmt.Sprintf("%x", md5.Sum([]byte(dir)))
	cacheFilePath := config.CurrentDir + "/cache/" + cacheFileName + ".json"
	var markdownLists MarkdownList
	cacheFile, cacheErr := ioutil.ReadFile(cacheFilePath)
	if cacheErr == nil && json.Unmarshal(cacheFile, &markdownLists) == nil {
		return markdownLists, nil
	}

	markdownLists, err := getMarkdownList(dir)
	if err != nil {
		return markdownLists, err
	}

	sort.Sort(markdownLists)
	markdownListJson, err := json.Marshal(markdownLists)
	if err != nil {
		return markdownLists, err
	}

	cacheDir := config.CurrentDir + "/cache"
	cacheInfo, err := os.Stat(cacheDir)
	if err != nil || !cacheInfo.IsDir() {
		if os.Mkdir(cacheDir, os.ModePerm) != nil {
			return markdownLists, err
		}
	}

	if err = ioutil.WriteFile(cacheFilePath, markdownListJson, os.ModePerm); err != nil {
		return markdownLists, err
	}

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
