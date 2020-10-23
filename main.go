package main

import (
	"bgm/utils"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/spf13/pflag"
	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v2"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type PostMeta struct {
	Alias string
	Title string
	Date  string
	Tags  []string
}

type SiteInfo struct {
	Name string
	url  string
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func readPost(filePath string) (PostMeta, string) {
	content, err := utils.ReadAll(filePath)
	if err != nil {
		// todo 404 : file not exist
		log.Fatal("read file err...", err)
	}
	lines := strings.Split(string(content), "\n")
	yamlStr := ""
	markdownStr := ""
	findSplitLine := false
	for _, line := range lines {
		if findSplitLine {
			markdownStr += line + "\n"
		} else {
			if strings.TrimSpace(strings.ReplaceAll(line, "-", "")) == "" && strings.Contains(line, "-") {
				log.Print("ok")
				findSplitLine = true
			} else {
				yamlStr += line + "\n"
			}
		}
	}
	return parsePost([]byte(yamlStr), []byte(markdownStr))
}

func parsePost(yamlStr []byte, markdownStr []byte) (PostMeta, string) {
	// 解析文章元信息
	meta := PostMeta{}
	err := yaml.Unmarshal(yamlStr, &meta)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 渲染 markdown
	var buf bytes.Buffer
	if err := goldmark.Convert(markdownStr, &buf); err != nil {
		panic(err)
	}

	return meta, buf.String()
}

func getPosts() []PostMeta {
	files, _, err := utils.GetFilesAndDirs("data/posts")
	if err != nil {
		// todo 404 : file not exist
		log.Fatal("can not read posts dir...", err)
	}
	filesCount := len(files)
	posts := make([]PostMeta, filesCount)
	for idx, postFile := range files {
		meta, _ := readPost(postFile)
		parts := strings.Split(postFile, "/")
		meta.Alias = strings.Replace(parts[len(parts)-1], ".md", "", 1)
		posts[idx] = meta
	}

	sort.Slice(posts, func(i, j int) bool {
		if posts[i].Date > posts[j].Date {
			return true
		}
		return false
	})
	return posts
}

func main() {
	posts := getPosts()
	switch os.Getenv("MODE") {
	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)
	case "DEBUG":
		gin.SetMode(gin.DebugMode)
	case "TEST":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	pflag.Parse()
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	section := cfg.Section("server")
	addr, _ := section.GetKey("addr")

	section = cfg.Section("site")
	siteName, _ := section.GetKey("name")
	siteUrl, _ := section.GetKey("url")
	site := SiteInfo{siteName.String(), siteUrl.String()}

	r := gin.Default()
	r.Static("/static", "static")
	r.Use(gin.Recovery())
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	if gin.IsDebugging() {
		r.HTMLRender = utils.NewDebug("templates")
	} else {
		r.HTMLRender = utils.NewProduction("templates")
	}

	r.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"site": site,
			"addr": addr.String(),
			"mode": os.Getenv("MODE"),
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.twig", utils.Context{
			"site":  site,
			"posts": posts,
		})
	})

	r.GET("/about", func(c *gin.Context) {
		meta, html := readPost("data/pages/about.md")
		c.HTML(http.StatusOK, "about.twig", utils.Context{
			"site": site,
			"meta": meta,
			"html": html,
		})
	})

	r.GET("/_refresh", func(c *gin.Context) {
		posts = getPosts()
		c.JSON(http.StatusOK, gin.H{
			"result": "OK",
		})
	})

	r.GET("/post/:alias", func(c *gin.Context) {
		alias := c.Param("alias")
		filePath := "./data/posts/" + alias + ".md"
		meta, html := readPost(filePath)
		c.HTML(http.StatusOK, "post.twig", utils.Context{
			"site": site,
			"meta": meta,
			"html": html,
		})
	})
	r.Run(addr.String())
}
