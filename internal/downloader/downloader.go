package downloader

type Downloader interface {
}

type downloader struct {
}

func NewDownloader() *downloader {
	return &downloader{}
}
