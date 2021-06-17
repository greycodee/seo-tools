package baidu

import (
	"errors"
	"os"
)
const baseUrl = "http://data.zz.baidu.com"
const submitUrlBatch = "/urls"

type IndexService struct {

}


func getBaiduEnv() (token string,siteUrl string,err error) {
	token = os.Getenv("BAIDU_TOKEN")
	if token == "" {
		return "","",errors.New("请设置 BAIDU_TOKEN 环境变量")
	}
	siteUrl = os.Getenv("SITE_URL")
	if siteUrl == "" {
		return "","",errors.New("请设置 SITE_URL 环境变量")
	}
	return
}