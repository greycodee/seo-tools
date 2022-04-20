package main

import (
	"flag"
	"github.com/greycodee/seo-tools/core"
	"github.com/greycodee/seo-tools/core/baidu"
	"github.com/greycodee/seo-tools/core/bing"
	"github.com/greycodee/seo-tools/core/google"
	"github.com/greycodee/seo-tools/tool/parse"
	"log"
	"regexp"
)

var (
	pushUrl string
	pushType string
	filePath string
	webMaster string
	pushCount int
)

func init() {

	flag.StringVar(&pushUrl,"u","","singe url publish\n(Used when the -t parameter is single)")

	flag.StringVar(&pushType,"t","single","push urls Type:\n	urlsFile\n	sitemap\n")
	flag.StringVar(&filePath,"f","","local file path or remote file path\n(Used when the -t parameter is urlsFile or sitemap)")

	flag.StringVar(&webMaster,"w","all","chose webmaster service:\n	google\n	baidu\n	bing\n")
	flag.IntVar(&pushCount,"n",0,"push count.If not configured, all are pushed by default. Only valid for sitemap")
	flag.Parse()
}
func main()  {
	start()
}

func start()  {
	switch pushType {
		case "single":
			singlePush()
			break
		case "urlsFile":
			urlsFilePush()
			break
		case "sitemap":
			sitemapPush()
			break
		default:
			panic("请选择正确的推送类型【urlsFile、sitemap】")
	}
}

func singlePush()  {
	// 单条 url
	if pushUrl == "" {
		panic("输入要推送的 URL 网址")
	}
	componentPushUrl(pushUrl)
}

func urlsFilePush()  {
	p := parse.NewSiteMap(pushCount)
	parseFileAndPush(p)
}

func sitemapPush()  {
	p := parse.NewSiteMap(pushCount)
	//var p = &parse.SiteMap{2}
	parseFileAndPush(p)
}

func parseFileAndPush(p parse.Parse)  {
	if filePath == "" {
		panic("请输入本地文件或远程文件路径")
	}
	isLocalFile := false
	// 判断文件是远程还是本地
	reg := regexp.MustCompile("[a-zA-z]+://[^\\s]*")
	if reg == nil { //解释失败，返回nil
		panic("正则表达式异常")
	}
	result:=reg.FindAllString(filePath,-1)
	if len(result) == 0 {
		// 说明不是 URL，默认认为它是本地文件路径
		isLocalFile = true
	}

	var urls []string
	var err error

	if isLocalFile {
		urls,err = p.ParseLocalFile(filePath)
		if err != nil {
			panic("解析本地文件失败")
		}
	}else {
		urls,err = p.ParseRemoteFile(filePath)
		if err != nil {
			panic("解析远程文件失败")
		}
	}
	componentPushUrlList(urls)
}

func chosePushWebMaster() (pushServices []core.PushUrl) {
	switch webMaster {
		case "google":
			pushServices = append(pushServices,new(google.IndexService))
			break
		case "baidu":
			pushServices = append(pushServices,new(baidu.IndexService))
			break
		case "bing":
			pushServices = append(pushServices,new(bing.IndexService))
			break
		case "all":
			pushServices = append(pushServices,new(google.IndexService))
			pushServices = append(pushServices,new(baidu.IndexService))
			pushServices = append(pushServices,new(bing.IndexService))
			break
		default:
			panic("WebMaster选择错误")
	}
	return
}

func componentPushUrl(url string) {
	pushServices := chosePushWebMaster()
	for _,v := range pushServices {
		resp, err := v.PushUrl(url)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("推送结果：%s",resp)
	}
}
func componentPushUrlList(urls []string) {
	pushServices := chosePushWebMaster()
	for _,v := range pushServices {
		resp, err := v.PushListUrl(urls)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("推送结果：%s",resp)
	}
}