# mdblog - Markdown Blog Engine 

> mdblog 是基于 go 语言开发的，适合用来学习和展示 markdown 文档的个人博客。

## 文章目录结构
```text
content       --------------- 博客所有文章所在目录
├── test      --------------- 文章test所在文件夹
│   ├── post.md  ------------ 文章test的正文内容
├── ├── meta.yml ------------ 文章test的meta信息
├── Works.md  --------------- 一级页面Work
├── About.md  --------------- 一级页面About

```

## meta.yml 
```yaml
tags:
    - blockchain
    - go
```

## config.yml 配置说明
- port: 服务端口 
- page_size: 首页每页显示文章数量
- description_len: 文章描述长度
- author: 作者
- webhook_secret: 博客更新钩子的秘钥，需要和仓库设置的秘钥一致
- githook_url: webHook地址
- utterances_repo: 是否开启utterances评论，留空则表示不开启评论，否则填写评论仓库name/repo(指向的仓库必须是公开并且可以被评论的，具体使用请访问 https://utteranc.es)
- document_path: 文章目录
- time_layout: 时间格式
- site_name: 网站的名字
- site_keywords: 网站关键字
- site_description: 网站描述
- category_doc_number: 每个分类的下面最多展示多少篇文章
- theme_color: 站点主题颜色
- theme_option: 站点主题可选颜色
- utterances_repo:

