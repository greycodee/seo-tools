## 环境变量设置
```shell
# 站点域名【当要推送 Baidu 或 Bing 时设置】
SITE_URL=http://blog.mjava.top

# Bing apiKey 配置【当要推送 Bing 时设置】
BING_API_KEY="xxxxxxx"

# Baidu Token 配置【当要推送 Baidu 时设置】
BAIDU_TOKEN="xxxxxxx"

# Google 密钥文件路径环境变量【当要推送 Google 时设置】
GOOGLE_APPLICATION_CREDENTIALS="/xxx/xxx/xxx.json"
```

## 命令使用

![命令使用图解](./flow.jpg)

### -t 选项
设置推送的类型，默认为 `single`,可选值有：
- single 单条 url 推送
- urlsFile 批量 url 文件推送

  > 文件格式(每一行一个 URL 地址):
  ```text
      http://example.com/index.html
      http://example.com/index.html
      http://example.com/index.html
  ```
- sitemap 站点地图推送
  > 站点地图 XML 格式
  ```xml
    <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
      <url>
          <loc>http://blog.mjava.top/gof/gof-singleton/</loc>
          <lastmod>2021-06-15T07:14:24.487Z</lastmod>
      </url>
     </urlset>
  ```

### -u 选项
设置要推送的 url，当 `-t` 选项的值为 `single` 时生效

示例：
```shell
./seo-tools -u https://example.com/test.html
```

### -f 选项
指定要推送当 urls 文件或站点地图文件地址

可以是本地文件或远程文件地址

当 `-t` 选项的值为 `urlsFile` 或 `sitemap` 时生效

示例：
```shell
./seo-tools -t sitemap -f https://example.com/sitemap.xml

# or

./seo-tools -t sitemap -f /root/home/sitemap.xml
```

### -w 选项
> 默认是全部推送

选择要推送当站长，可选项：
- baidu
- bing
- google

示例：
```shell
# 只推送到百度站长
./seo-tools -u https://example.com/test.html -w baidu

# 只推送到必应站长
./seo-tools -u https://example.com/test.html -w bing

# 只推送到谷歌站长
./seo-tools -u https://example.com/test.html -w google
```
