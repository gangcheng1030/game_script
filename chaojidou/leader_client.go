package chaojidou

import (
	"net/http"
	"net/url"
)

func Register(remoteAddr string, localAddr string) error {
	url1 := "http://" + remoteAddr + "/leader"

	urlValues := url.Values{}
	urlValues.Set("addr", localAddr)

	_, err := http.PostForm(url1, urlValues)
	return err
}
