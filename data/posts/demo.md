title: Demo
date: 2020-10-19
---

### 无序列表

- 张三
- 李四
- 王五

### 有序列表：
1. 法国
2. 意大利
3. 德国

### 代码

```
# 筛选 apache2 的进程
ps aux | grep httpd

# 停止 apache2
sudo apachectl -k stop

# 关闭 apache2 跟随系统启动
sudo launchctl unload -w /System/Library/LaunchDaemons/org.apache.httpd.plist

# 启动使用 brew 安装的 Nginx
sudo brew services start nginx
```

### 图片

![](/static/images/avatar.jpg)

### 未来计划

目前仅支持原生 [Markdown](https://tools.ietf.org/html/rfc7763)，暂时不支持表格。

未来的目标包括让 BGM 更加灵活、更加强大，当然，会继续保持开箱即用的简洁风格。