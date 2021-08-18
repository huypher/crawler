package downloader

type HttpDownloader interface {
}

type httpDownloader struct {
}

func NewHttpDownloader() *httpDownloader {
	return &httpDownloader{}
}
