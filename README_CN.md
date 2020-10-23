# BGM

BGM 是指使用 [Gin Web Framework](https://github.com/gin-gonic/gin) 和 [Markdown](https://tools.ietf.org/html/rfc7763) 开发的博客系统。

### 下载&启动

从 [Github 的 BGM Release 页面](https://github.com/lucups/bgm/releases/) 下载最新版本的二进制文件，解压，即可开始运行。

运行命令如下：

```
./bgm
```

如果要放到后台运行，可以使用 `nohup` 命令:

```
nohup ./bgm > bgm.log &
```

如果要停止进程，可以使用 `ps` 进程找到进程 ID，然后通过 `kill` 命令杀掉即可：

```
ps aux | grep bgm

kill -9 进程ID
```

### 内容编辑

为了简化配置，暂时默认所有的内容文件都在项目的 `data` 目录下:

```
data/posts # 存放文章，文件名会自动匹配路由，
data/pages # 存放单页面，目前只有 about.md
```

注意：文件名后缀必须为 `.md`。

关于内容格式，分为两部分，以第一个全部为 `-` 符号的行作为分割，前面是文章的元信息，采用 YAML 格式；
后面是文章的内容，采用 Markdown 格式。

文章的元信息包括：`标题(title)`、`日期(date)`、`标签(tags，数组格式)`，其中标签后续支持渲染，当然后续也会增加更多的元信息属性。

示例：

```
title: Hello World
date: 2020-10-10
tags:
    - hello
    - go
------------------

Hello everyone!

[BGM](https://github.com/lucups/bgm) 是一个快速搭建简单博客的工具。

它基于 [Gin](https://github.com/gin-gonic/gin) 开发，使用 [Markdown](https://tools.ietf.org/html/rfc7763) 作为内容的格式。

您只需要简单的启动它，并开始编写 Markdown 格式的文本，即可看到渲染效果，方便快捷。

欢迎使用！
```

### 内容刷新

为了加快访问速度，文章列表只会在程序启动时加载一次，后面直接就放在内存里，所以如果新增文章或修改了文章的元信息，需要调用刷新接口。

刷新接口路径为 `/_refresh`，在本地即 `http://localhost:9600/_refresh` 。

刷新后，新文章或修改的元信息才会生效。

文章的内容部分，每次都是直接读取，所以无需刷新，实时生效。

### 模板定制

当然，您可以随心所欲地修改原生的模板文件，使之达到您想要的效果。


### 项目依赖

- [Gin Web Framework](https://github.com/gin-gonic/gin) A web framework written in Go. 
- [Pongo2](https://github.com/flosch/pongo2) A Django-syntax like templating-language. (Also like Jinja/Twig/Nunjucks)
- [goldmark](https://github.com/yuin/goldmark) 解析 Markdown。
- [ini](github.com/go-ini/ini) 解析 `ini` 文件。
- [yaml](gopkg.in/yaml.v2) 解析 `YAML`。
