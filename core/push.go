package core

type PushUrl interface {
	PushUrl(url string) (resp string,err error)
	PushListUrl(urlList []string) (resp string,err error)
}
