package parse

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)
/*
	站点地图文件解析
	站点地图 XML 格式
	<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
		<url>
			<loc>http://blog.mjava.top/gof/gof-singleton/</loc>
			<lastmod>2021-06-15T07:14:24.487Z</lastmod>
		</url>
	</urlset>
*/
type SiteMap struct {
	urlCount int
}

func NewSiteMap(pushCount int) *SiteMap {
	return &SiteMap{pushCount}
}

type urlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Url []urlItem    `xml:"url"`
}
type urlItem struct {
	XMLName xml.Name `xml:"url"`
	Loc string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

func (s SiteMap) ParseRemoteFile(url string) (urls []string,err error) {
	log.Printf("开始解析远程sitemap：%s",url)
	resp,err:=http.Get(url)
	if err != nil{
		return
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return s.parseXml(content)
}

func (s SiteMap) ParseLocalFile(filePath string) (urls []string,err error) {
	log.Printf("开始解析本地sitemap：%s",filePath)
	file,err := os.Open(filePath)
	if err != nil{
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	content,err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	return s.parseXml(content)
}


func (s SiteMap) parseXml(content []byte)  (urls []string,err error) {
	log.Println("开始解析XML")
	sm := urlSet{}
	err = xml.Unmarshal(content, &sm)
	if err != nil {
		return
	}
	if s.urlCount>0 {
		sm.Url = sm.Url[:s.urlCount]
	}
	for _,v := range sm.Url{
		log.Printf("【SITE-MAP】解析到url：%s",v.Loc)
		urls = append(urls,v.Loc)
	}
	return
}