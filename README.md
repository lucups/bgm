# BGM

BGM means a blog program build by [Gin Web Framework](https://github.com/gin-gonic/gin) and [Markdown](https://tools.ietf.org/html/rfc7763).

- [中文版](README_CN.md)

### Download & Start

Download the latest release binary file from [Github 的 BGM Release 页面](https://github.com/lucups/bgm/releases/), and run it by this command:

```
./bgm
```

If you want to make it background，you can use `nohup` command:

```
nohup ./bgm > bgm.log &
```

Kill it by `ps` and `kill` commands:

```
ps aux | grep bgm

kill -9 processId
```

### Edit & Publish

For simple, all posts should be put in `data` dir:

```
data/posts # Put your articles here.
data/pages # Put single pages, now there is only one page: about.md
```

Tips：the suffix of all files should be `.md`。

The post/page's content will split two parts, 

As for the content format, it is divided into two parts.

The first line is all '-' symbol as the segmentation. The front is the meta information of the article, which adopts yaml format;

The following is the content of the article in markdown format.

The meta information of the article includes: `title` ,`date`,`tags (array format)`,
 in which the subsequent tags support rendering, of course, more meta information attributes will be added later.

A simple example：

```
title: Hello World
date: 2020-10-10
tags:
    - hello
    - go
------------------

Hello everyone!

[BGM](https://github.com/lucups/bgm) is a quick tool to build a simple blog.

It's based on [Gin](https://github.com/gin-gonic/gin), and use [Markdown](https://tools.ietf.org/html/rfc7763) as the format of the content  .

You just need to simply start it, and start writing markdown format text, you can see the rendering effect, convenient and fast.
Welcome to use it！
```
###Content refresh
In order to speed up the access, the article list will be loaded only once when the program starts, and then it will be directly put into the memory. Therefore, if you add new articles or modify the meta information of articles, you need to call the refresh interface.

The refresh interface path is`/_ Refresh ', which means "Refresh"` http://localhost :9600/_ refresh` 。

The new article or modified meta information will not take effect until it is refreshed.

The content of the article is read directly each time, so it does not need to be refreshed and takes effect in real time.

### Template customization

Of course, you can modify the native template file to achieve the desired effect.

### Dependencies

- [Gin Web Framework](https://github.com/gin-gonic/gin) A web framework written in Go. 
- [Pongo2](https://github.com/flosch/pongo2) A Django-syntax like templating-language. (Also like Jinja/Twig/Nunjucks)
- [goldmark](https://github.com/yuin/goldmark) Parse Markdown。
- [ini](github.com/go-ini/ini) Parse `ini`。
- [yaml](gopkg.in/yaml.v2) Parse `YAML`。

