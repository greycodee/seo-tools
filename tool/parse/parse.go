package parse

type Parse interface {
	ParseRemoteFile(url string) (urls []string,err error)
	ParseLocalFile(path string) (urls []string,err error)
}
