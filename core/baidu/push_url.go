package baidu

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func (b IndexService) PushUrl(url string) (string,error) {
	return push([]byte(url))
}

func (b IndexService) PushListUrl(urlList []string) (string,error) {
	var urls string
	for _,v := range urlList {
		fmt.Println(v)
		urls += v
		urls += "\n"
	}
	return push([]byte(urls))
}

func push(body []byte) (string,error) {
	log.Println("【Baidu】开始推送")
	token,siteUrl,err := getBaiduEnv()
	if err != nil {
		return "",err
	}
	requestUrl := baseUrl + submitUrlBatch +"?site="+siteUrl+"&token="+token
	response,err := http.Post(requestUrl,"text/plain",bytes.NewReader(body))
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
