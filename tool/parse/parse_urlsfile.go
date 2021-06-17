package parse

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

/*
	批量 URL 文件解析
	文件格式:
		http://example.com/index.html
		http://example.com/index.html
		http://example.com/index.html
	每一行一个 URL 地址
*/
type UrlsFile struct {

}

func (u UrlsFile) ParseRemoteFile(url string) (urls []string,err error) {
	log.Printf("开始解析远程urlfile：%s",url)
	resp,err:=http.Get(url)
	if err != nil{
		return
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return parseUrlsFile(resp.Body)
}

func (u UrlsFile) ParseLocalFile(path string) (urls []string,err error) {
	f,err := os.Open(path)
	if err != nil{
		return
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return parseUrlsFile(f)
}

func parseUrlsFile(data io.Reader) (urls []string,err error) {
	log.Println("开始解析 URL 文件")
	s := bufio.NewScanner(data)
	for s.Scan() {
		log.Printf("【URL-FILE】解析到url：%s",s.Text())
		urls = append(urls,s.Text())
	}
	return
}