package google

import (
	"context"
	"encoding/json"
	"google.golang.org/api/indexing/v3"
	"log"
	"sync"
)

func (b IndexService) PushUrl(url string) (string,error) {
	err := getGoogleEnv()
	if err != nil {
		return "",err
	}
	return push(update,url)
}

func (b IndexService) PushListUrl(urlList []string) (string,error) {
	err := getGoogleEnv()
	if err != nil {
		return "",err
	}

	var res result
	var wg sync.WaitGroup
	for _,v := range urlList{
		wg.Add(1)
		go syncPush(v,&res,&wg)
	}
	wg.Wait()
	r,_:=json.Marshal(res)
	return string(r),nil
}

func syncPush(url string,res *result,wg *sync.WaitGroup)  {
	log.Printf("【Google】开始处理：%s",url)
	r,err := push(update,url)
	// 添加响应结果
	item := rows{
		Url:   url,
		Resp:  r,
		Error: err,
	}
	res.Rows = append(res.Rows, item)
	wg.Done()
}

func push(pushType string,url string) (string,error)  {
	log.Printf("【Google】开始推送")
	ctx := context.Background()
	indexService,err:=indexing.NewService(ctx)
	if err != nil {
		return "", err
	}
	publish:=indexService.UrlNotifications.Publish(&indexing.UrlNotification{
		Type:            pushType,
		Url:            url,
	})
	res,err:=publish.Do()
	if err != nil {
		log.Fatal(err)
	}
	by,err:=res.MarshalJSON()
	if err != nil {
		return "", err
	}
	return string(by),nil
}

const (
	none = "URL_NOTIFICATION_TYPE_UNSPECIFIED"
	update = "URL_UPDATED"
	deleted = "URL_DELETED"

)

type result struct {
	Rows []rows `json:"rows"`
}
type rows struct {
	Url string `json:"url"`
	Resp string `json:"resp"`
	Error error `json:"error"`
}

