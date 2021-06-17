package bing

import (
	"errors"
	"os"
)

const baseUrl = "https://ssl.bing.com"
const submitUrl = "/webmaster/api.svc/json/SubmitUrl"
const submitUrlBatch = "/webmaster/api.svc/json/SubmitUrlBatch"

type IndexService struct {
}

func getBingEnv() (apiKey string,siteUrl string,err error) {
	apiKey = os.Getenv("BING_API_KEY")
	if apiKey == "" {
		return "","",errors.New("请设置 BING_API_KEY 环境变量")
	}
	siteUrl = os.Getenv("SITE_URL")
	if siteUrl == "" {
		return "","",errors.New("请设置 SITE_URL 环境变量")
	}
	return
}