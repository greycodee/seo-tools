package bing

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type requestBody struct {
	// 站点域名地址
	SiteUrl string `json:"siteUrl"`
	// 要提交的 url
	Url string `json:"url"`
	// 批量提交的 url 列表
	UrlList []string `json:"urlList"`
}

func (b IndexService) PushUrl(url string) (string,error) {
	apiKey,siteUrl,err := getBingEnv()
	if err != nil {
		return "",err
	}
	requestUrl := baseUrl + submitUrl +"?apikey="+apiKey
	req,_ := json.Marshal(requestBody{
		SiteUrl: siteUrl,
		Url:     url,
	})
	return push(requestUrl,req)
}

func (b IndexService) PushListUrl(urlList []string) (string,error) {
	apiKey,siteUrl,err := getBingEnv()
	if err != nil {
		return "",err
	}
	requestUrl := baseUrl + submitUrlBatch +"?apikey="+apiKey
	req,_ := json.Marshal(requestBody{
		SiteUrl: 	 siteUrl,
		UrlList:     urlList,
	})
	return push(requestUrl,req)
}

func push(url string,body []byte) (string,error) {
	log.Printf("【Bing】开始推送URL:%s",url)
	response,err := http.Post(url,"application/json",bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(content),nil
}


